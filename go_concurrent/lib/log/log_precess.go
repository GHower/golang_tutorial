package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 读取者
type Reader interface {
	Read(rc chan []byte)
}

// 写出
type Writer interface {
	Write(wc chan *Message)
}

type LogProcess struct {
	// 各模块传递数据
	rc    chan []byte
	wc    chan *Message
	read  Reader
	write Writer
}

type ReadFromFile struct {
	path string
}
type WriteToInfluxDB struct {
	influxDBDsn string
}

// Message 日志消息的结构
type Message struct {
	TimeLocal                    time.Time
	ByteSent                     int
	Path, Method, Scheme, Status string
	UpstreamTime, RequestTime    float64
}

// SystemInfo 系统状态监控结构体
type SystemInfo struct {
	HandleLine   int     `json:"handleLine"`   // 总处理日志行数
	Tps          float64 `json:"tps"`          // 系统吞吐量
	ReadChanLen  int     `json:"readChanLen"`  // read channel长度
	WriteChanLen int     `json:"writeChanLen"` // write channel 长度
	RunTime      string  `json:"runTime"`      // 运行总时间
	ErrNum       int     `json:"errNum"`       // 错误数
}

// 定义常量，分别表示行数和错误数
const (
	TypeHandleLine = 0
	TypeErrNum     = 1
)

// 监控者处理的管道
var TypeMonitorChan = make(chan int, 200)

// Monitor 监控者
type Monitor struct {
	startTime time.Time
	data      SystemInfo
	tpsSli    []int
}

func (m *Monitor) start(lp *LogProcess) {
	go func() {
		for n := range TypeMonitorChan {
			switch n {
			case TypeHandleLine:
				m.data.HandleLine += 1
			case TypeErrNum:
				m.data.ErrNum += 1

			}
		}
	}()
	// 5秒吞吐量
	ticker := time.NewTimer(time.Second * 5)
	go func() {
		<-ticker.C
		m.tpsSli = append(m.tpsSli, m.data.HandleLine)
		if len(m.tpsSli) > 2 {
			m.tpsSli = m.tpsSli[1:]
		}
	}()
	http.HandleFunc("/monitor", func(writer http.ResponseWriter, request *http.Request) {
		m.data.RunTime = time.Now().Sub(m.startTime).String()
		m.data.ReadChanLen = len(lp.rc)
		m.data.WriteChanLen = len(lp.wc)
		if len(m.tpsSli) >= 2 {
			m.data.Tps = float64(m.tpsSli[1]-m.tpsSli[0]) / 5
		}
		ret, _ := json.MarshalIndent(m.data, "", "\t")
		io.WriteString(writer, string(ret))

	})
	http.ListenAndServe(":9193", nil)
}
func (r *ReadFromFile) Read(rc chan []byte) {
	// 读取模块
	//line := "message"
	//rc <- line
	f, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("文件打开失败:%s", err.Error()))
	}
	// 从文件末尾开始逐行读取offset为偏移量,whence=2为文件末尾
	f.Seek(0, 2)
	rd := bufio.NewReader(f)
	for {
		// 读一行
		line, err := rd.ReadBytes('\n')
		// 文件末尾
		if err == io.EOF {
			TypeMonitorChan <- TypeHandleLine

			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("读取出错：%s", err.Error()))
		}
		TypeMonitorChan <- TypeHandleLine
		rc <- line[:len(line)-1]
	}

}
func (w *WriteToInfluxDB) Write(wc chan *Message) {
	// influx数据库写入，这里不写

	// 简单写入
	for v := range wc {
		fmt.Print(v)
	}
}

func (l *LogProcess) Process() {
	// 解析
	/**
	172.0.0.12 - -[04/Mar/2018:13:49:52 +0000] http "GET/foo?query=t HTTP/1.0" 200 2133 "-" "KeepAliveclient" "-" 1.005 1.854

	([\d\.]+)\s+([^ \[]+]\s([^ \[]+)\s+\[([^\]]+)\]\s+([a-z]+)\s+\"([^"]+)\"\s+(\d{3})\s+(\d+)\s+\"([^"]+)\"\s+\"(.*?)\"\s+\"([\d\.-]+)\"\s+([\d\.-]+)\s+([\d\.-]+)
	*/

	r := regexp.MustCompile(`([\d\.]+)\s+([^ \[]+]\s([^ \[]+)\s+\[([^\]]+)\]\s+([a-z]+)\s+\"([^"]+)\"\s+(\d{3})\s+(\d+)\s+\"([^"]+)\"\s+\"(.*?)\"\s+\"([\d\.-]+)\"\s+([\d\.-]+)\s+([\d\.-]+)`)

	loc, _ := time.LoadLocation("Asia/Shanghai")
	for v := range l.rc {
		ret := r.FindStringSubmatch(string(v))
		if len(ret) != 14 {
			TypeMonitorChan <- TypeErrNum

			log.Println("匹配失败:", string(v))
			continue
		}
		message := &Message{}
		t, err := time.ParseInLocation("04/Jan/2006:15:04:05 +0000", ret[4], loc)
		if err != nil {
			TypeMonitorChan <- TypeErrNum
			log.Println("解析异常:", err.Error(), ret[4])
			continue
		}
		message.TimeLocal = t
		//
		byteSend, _ := strconv.Atoi(ret[8])
		message.ByteSent = byteSend

		//
		reqSli := strings.Split(ret[6], " ")
		if len(reqSli) != 3 {
			TypeMonitorChan <- TypeErrNum
			log.Println("strings异常：", ret[6])
			continue
		}
		message.Method = reqSli[0]
		u, err := url.Parse(reqSli[1])
		if err != nil {
			log.Println("url有错误:", err.Error())
			TypeMonitorChan <- TypeErrNum
			continue
		}
		message.Path = u.Path

		message.Scheme = ret[5]
		message.Status = ret[7]

		upstreamTime, _ := strconv.ParseFloat(ret[12], 64)
		requestTime, _ := strconv.ParseFloat(ret[13], 64)
		message.UpstreamTime = upstreamTime
		message.RequestTime = requestTime

		//println(strings.ToUpper(string(v)))
		//l.wc <- strings.ToUpper(string(v))
	}
}

func main() {
	// 读取的日志文件和influxDsn信息
	var path, influxDsn string
	flag.StringVar(&path, "path", "./access.log", "read file path")
	flag.StringVar(&influxDsn, "influxDsn", "http://127.0.0.1:8086@username@password@dbname@timestamp", "")
	flag.Parse()

	r := &ReadFromFile{
		path: path,
	}
	w := &WriteToInfluxDB{
		influxDBDsn: influxDsn,
	}

	lp := &LogProcess{
		rc:    make(chan []byte),
		wc:    make(chan *Message),
		read:  r,
		write: w,
	}

	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)

	// 启用自己的monitor
	m := &Monitor{
		startTime: time.Now(),
		data:      SystemInfo{},
	}
	m.start(lp)

	//time.Sleep(30 * time.Second)
}
