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
	// 循环读取服务器输入
	go func() {
		for {
			var msg []byte = make([]byte, 26)
			length, _ := conn.Read(msg)
			if length > 0 {
				print("<", length, ":")
				for i := 0; ; i++ {
					if msg[i] == 0 {
						break
					}
					fmt.Printf("%c", msg[i])
				}
				print(">")
			}
		}
	}()
	for {
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClien + " says:" + trimmedInput))
	}

}
