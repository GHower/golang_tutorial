package main

import (
	"fmt"
	"time"
)

func main() {
	f1()
	f1()
	time.Sleep(time.Second)
}
func f1() {
	a := 1
	go func() {
		a = 2
		fmt.Println(a)
	}()
}
