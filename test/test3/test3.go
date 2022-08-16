package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		func() {
			fmt.Println(i)
			defer mclose()
		}()
	}
}
func mclose() {
	fmt.Println("close!")
}
