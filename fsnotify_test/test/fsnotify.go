package main

import (
	"bufio"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/lovejoy/seelog"
	"io"
	"log"
	"os"
	"time"
)

func Init(conf ...string) error {
	logger, err := seelog.LoggerFromConfigAsFile("fsnotify_test\\test\\log.xml")
	if err != nil {
		seelog.Criticalf("parsing log config %s failed: %s", "log.xml", err)
		return err
	}
	seelog.ReplaceLogger(logger)
	seelog.Infof("init log from %s success", "log.xml")
	seelog.Flush()
	return nil
}

func main() {
	logfile := "D:\\goland_projects\\golang_tutorial\\fsnotify_test\\test"
	Init()
	// Start listening for events.
	go watchHandle(logfile)

	//go genLog()
	// Block main goroutine forever.
	<-make(chan struct{})
}

func genLog() {
	for i := 1; i < 1000; i++ {
		seelog.Infof("test log info  %d", i)
		//seelog.Flush()
		time.Sleep(time.Millisecond * 1000)
	}
}
func genLog2() {
	for i := 1; i < 1000; i++ {
		seelog.Infof("test log info  %d", i)
		//seelog.Flush()
		time.Sleep(time.Millisecond * 1000)
	}
}
func watchHandle(filename string) error {
	// 创建监听器
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		seelog.Errorf("watcher error: %d", err)
	}
	defer watcher.Close()

	// 添加监听路径
	err = watcher.Add(filename)
	if err != nil {
		log.Fatal(err)
	}
	// 添加监听路径 2
	err = watcher.Add("D:\\goland_projects\\golang_tutorial\\fsnotify_test\\tail_test")
	if err != nil {
		log.Fatal(err)
	}
	// 打开文件
	file, err := os.Open("D:\\goland_projects\\golang_tutorial\\fsnotify_test\\tail_test" + "\\server.log")
	if err != nil {
		fmt.Printf("open file error: %s \n", err)
	}
	r := bufio.NewReader(file)

	_, _ = r.ReadBytes('\n')
	for {
		by, err := r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			fmt.Printf("other error: %s \n", err)
			return err
		}
		fmt.Println(string(by))

		//if err != io.EOF {
		//	continue
		//}
		if err = waitForChange(watcher); err != nil {
			fmt.Printf("watch error: %s \n", err)
			return nil
		}
	}

}
func waitForChange(watcher *fsnotify.Watcher) error {
	for {
		select {
		case event := <-watcher.Events:
			log.Println("event:", event)
			if event.Op == fsnotify.Write {
				log.Println("modified file:", event.Name)
				return nil
			}
		case err := <-watcher.Errors:
			log.Println("err:", err)
			return err
		}
	}
}
