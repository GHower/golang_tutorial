package main

import "fmt"
import "time"

func main() {
	c := make(chan int, 5)
	go func() {
		time.Sleep(15 * time.Second)
		x := <-c
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)
}

/* Output:
sending 10
(15 s later):
received 10
sent 10
*/
