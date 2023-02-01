package main

import (
	"bufio"
	"fmt"
	"github.com/nxadm/tail"
	"github.com/nxadm/tail/ratelimiter"
	"github.com/panjf2000/ants/v2"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime/debug"
	"runtime/pprof"
	"sync"
	"sync/atomic"
	"time"
)

var count = int64(0)

func main() {
	//preCreate()
	test2()
}

var threadProfile = pprof.Lookup("threadcreate")

type cTask struct {
	filepath string
}

var ioOpLimiter = ratelimiter.NewLeakyBucket(60, time.Second)

func test1() {
	for i := 0; i < 120; i++ {
		if i > 60 {
			time.Sleep(time.Second)
		}
		pour := ioOpLimiter.Pour(1)
		log.Printf("%d:%v:%v", i, pour, ioOpLimiter.Fill)
	}

}

func test2() {
	debug.SetMaxThreads(100)
	wg := sync.WaitGroup{}

	dirs, err := ioutil.ReadDir("D:\\proTemp\\")
	if err != nil {
		return
	}
	st := threadProfile.Count()
	//fmt.Printf("threads in starting: %d\n", threadProfile.Count())
	wg.Add(1)
	p, err := ants.NewPoolWithFunc(100, createTail)
	if err != nil {
		return
	}
	for _, f := range dirs {
		//if i%10 == 0 {
		//	time.Sleep(time.Millisecond)
		//}
		fp := fmt.Sprintf("D:\\proTemp\\%s", f.Name())
		p.Invoke(cTask{filepath: fp})
	}
	go func() {
		time.Sleep(time.Second * 10)
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("start: %d --- end: %d\n", st, threadProfile.Count())
}
func createTail(data interface{}) {
	task := data.(cTask)
	config := tail.Config{
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: io.SeekEnd,
		},
		ReOpen:    true,
		MustExist: false,
		Poll:      false,
		Follow:    true,
	}

	tails, err := tail.TailFile(task.filepath, config)
	if err != nil {
		log.Fatalf("%s err:%s", tails.Filename, err)
		return
	}
	atomic.AddInt64(&count, 1)
	fmt.Printf("threads in starting: %d 当前第:%d\n", threadProfile.Count(), count)

	go handler(tails)

	//log.Println(tails.Filename)
}

func handler(tails *tail.Tail) {
	for {
		select {
		case line, isOpen := <-tails.Lines:
			if !isOpen {
				log.Fatalf("gg:%s", tails.Filename)
			}
			log.Printf("=== text:%s", line.Text)
		}
	}
}
func preCreate() {
	for i := 1; i < 200000; i++ {
		fp := fmt.Sprintf("D:\\proTemp\\test%d.log", i)
		f, err := os.Create(fp)
		if err != nil {
			log.Println(err)
			return
		}
		writer := bufio.NewWriter(f)
		writer.WriteString("test")
		writer.Flush()
		f.Close()
	}
}
