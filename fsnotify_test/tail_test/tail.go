package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"github.com/lovejoy/seelog"
	"io"
	"time"
)

var stop chan struct{}

func main() {
	Init()
	// 持续生成日志
	go genLog()
	//go genLog()

	// tail监听
	//tailHandle()
	<-stop
}

func tailHandle() {
	fileName := "D:\\goland_projects\\golang_tutorial\\fsnotify_test\\tail_test\\server.log"

	//fileName := "./server.log"
	config := tail.Config{
		ReOpen:    true,                                          // 重新打开
		Follow:    true,                                          // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: io.SeekEnd}, // 从文件的哪个地方开始读
		MustExist: false,                                         // 文件不存在不报错
		Poll:      true,                                          // 轮询模式，windows下监听模式可能不产生事件,建议用poll
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	var (
		line *tail.Line
		ok   bool
	)
	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, line:%v\n", line)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line:", line.Text)
	}
}
func Init(conf ...string) error {
	logger, err := seelog.LoggerFromConfigAsFile("fsnotify_test\\tail_test\\log.xml")
	if err != nil {
		seelog.Criticalf("parsing log config %s failed: %s", "log.xml", err)
		return err
	}
	seelog.ReplaceLogger(logger)
	seelog.Infof("init log from %s success", "log.xml")
	seelog.Flush()
	return nil
}
func genLog() {
	for i := 1; i < 1000; i++ {
		seelog.Infof("test log info  %d", i)
		//seelog.Flush()
		time.Sleep(time.Second)
	}
}
