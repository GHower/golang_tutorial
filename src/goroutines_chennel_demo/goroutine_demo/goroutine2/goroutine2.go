package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(time.Second)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch)
}

func getData(ch chan string) {
	var input string
	// time.Sleep(1e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}

// Washington Tripoli London Beijing Tokio
