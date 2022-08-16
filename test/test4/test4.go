package main

import "time"

func main() {
	go println("hello")
	time.Sleep(time.Microsecond)
}

//func main() {
//	go println("hello")
//	runtime.Gosched()
//}
