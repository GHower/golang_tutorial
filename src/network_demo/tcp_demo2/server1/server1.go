package main

import (
	"fmt"
	"net"
)

func main() {
	// 网络包，监听 localhost:8080 作为服务器端
	listener, err := net.Listen("tcp", "localhost:8080")

	// 监听异常判断
	if err != nil {
		fmt.Println("监听异常：", err.Error())
		return
	}
	for {
		// 监听等待客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接收异常：", err.Error())
			return
		}
		// 用协程处理链接
		go doServerStuff(conn)
	}

}

func doServerStuff(conn net.Conn) {
	// 循环监听连接中发送的请求
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取异常：", err.Error())
			return
		}
		fmt.Printf("接收的数据:%v\n", string(buf[:len]))
	}
}
