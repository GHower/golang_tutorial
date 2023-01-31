package main

import (
	"golang_tutorial/gin_demo2/internal"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	_ = internal.InitModules("./gin_demo2/conf/dev/", []string{"app.yml"})
	defer internal.Destroy()

	//router.HttpServerRun()
	// main阻塞
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	//router.HttpServerStop()
}
