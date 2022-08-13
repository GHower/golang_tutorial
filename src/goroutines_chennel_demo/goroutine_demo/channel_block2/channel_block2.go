package main

import (
	"fmt"
	"time"
)

var cnt = 0

func main() {
	ch1 := make(chan int)
	go pump(ch1) // pump hangs
	go suck(ch1)
	//fmt.Println(<-ch1) // prints only 0

	time.Sleep(time.Second)
	fmt.Println(cnt)
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
		cnt++
	}
}
func suck(ch chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}
