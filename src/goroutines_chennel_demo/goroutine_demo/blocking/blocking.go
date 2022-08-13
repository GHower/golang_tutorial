package main

// blocking.go
// throw: all goroutines are asleep - deadlock!

import (
	"fmt"
	"time"
)

func f1(in chan int) {
	fmt.Println(<-in)
	fmt.Println(<-in)
}

func main() {
	out := make(chan int, 2)
	//out := make(chan int, 1) // solution 2
	// go f1(out)  // solution 1
	out <- 2
	out <- 2
	go f1(out)
	time.Sleep(time.Second)
}
