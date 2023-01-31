package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var HttpServerHandle *http.Server
var addrIp = "127.0.0.1"
var port = 9000

// HttpServerRun
// fixme: 所有配置从配置文件读取
func HttpServerRun() {
	gin.SetMode("debug")

	r := InitRouter()
	addr := fmt.Sprintf("%s:%d", addrIp, port)
	// 起一个handler，监听http
	HttpServerHandle = &http.Server{
		Addr:           addr,
		Handler:        r,
		ReadTimeout:    time.Duration(20) * time.Second,
		WriteTimeout:   time.Duration(20) * time.Second,
		MaxHeaderBytes: 1 << uint(20),
	}
	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n", addr)
		if err := HttpServerHandle.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", addr, err)
		}
	}()
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpServerHandle.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
