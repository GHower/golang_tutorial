package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 一个拨号链接，socket
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	inputReader := bufio.NewReader(os.Stdin)
	// 客户端名称
	clientName, _ := inputReader.ReadString('\n')

	// 去除尾端的EOF
	trimmedClien := strings.Trim(clientName, "\r\n")

	for {
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClien + " says:" + trimmedInput))
	}

}
