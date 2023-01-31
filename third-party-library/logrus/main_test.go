package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"testing"
)

// 定义等级
func Test_main1(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetReportCaller(true)

	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	logrus.Fatal("fatal msg")
	logrus.Panic("panic msg")
}

// 输出带上方法和文件名
func Test_main2(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)

	logrus.SetReportCaller(true)

	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	logrus.Fatal("fatal msg")
	logrus.Panic("panic msg")
}

// 携带字段
func Test_main3(t *testing.T) {
	logrus.WithFields(logrus.Fields{
		"name": "dj",
		"age":  18,
	}).Info("info msg")
}

// 重定向输出
func Test_main4(t *testing.T) {
	// 定向到结构体
	writer1 := &bytes.Buffer{}
	// 到标准输出
	writer2 := os.Stdout
	// 到文件
	writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}

	logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))
	logrus.Info("info msg")

	fmt.Println(writer1)
}

// 自定义
func Test_main5(t *testing.T) {
	log := logrus.New()

	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.JSONFormatter{})

	log.Info("info msg")
}

// 设置钩子
type AppHook struct {
	AppName string
}

func (h *AppHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *AppHook) Fire(entry *logrus.Entry) error {
	entry.Data["app"] = h.AppName
	return nil
}

func Test_main6(t *testing.T) {
	h := &AppHook{AppName: "awesome-web"}
	logrus.AddHook(h)

	logrus.Info("info msg")
}
