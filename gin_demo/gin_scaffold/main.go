package main

import (
	"gin_demo/gin_scaffold/router"
	"github.com/e421083458/golang_common/lib"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 设置链路监听
	lib.InitModule("./gin_scaffold/conf/dev/", []string{"base", "mysql", "postgresql"})
	defer lib.Destroy()

	// 尝试用通道限制流量为100,这是全局限制
	//limitChan := make(chan int, 100)
	//defer close(limitChan)
	// 启动http服务
	router.HttpServerRun()

	// 退出通道，通道接收到信号时，执行关闭操作
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	// 由于通道没有数据，这会阻塞
	<-quit
	router.HttpServerStop()
}
