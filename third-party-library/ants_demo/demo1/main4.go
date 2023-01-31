package main

import (
	"fmt"
	"sync"
)

func wrapper2(i int, wg *sync.WaitGroup) func() {
	return func() {
		fmt.Printf("hello from task:%d\n", i)
		if i%2 == 0 {
			panic(fmt.Sprintf("**panic** from task:%d", i))
		}
		wg.Done()
	}
}
