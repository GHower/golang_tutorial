package test

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestExp(t *testing.T) {
	jobChan := make(chan int64)
	defer close(jobChan)
	resultChan := make(chan []int64)
	defer close(resultChan)
	n := 100
	//wg := sync.WaitGroup{}
	// 1. 生成int64随机数共n个
	//wg.Add(1)
	go func(n int) {
		for i := 0; i < n; i++ {
			jobChan <- rand.Int63()
		}
	}(n)
	// 2. 开启24个协程从jobChan中取出随机数计算各位数的和，结果发送到resultChan
	//wg.Add(24)
	for i := 0; i < 24; i++ {
		go func() {
			for {
				sum := int64(0)
				x, ok := <-jobChan
				if !ok {
					break
				}
				for x := x; x != 0; x /= 10 {
					sum += x % 10
				}
				resultChan <- []int64{x, sum}
			}
		}()
	}
	// 3. 从resultChan取出结果并打印
	for {
		x := <-resultChan
		fmt.Println(x[0], x[1])
		if n--; n == 0 {
			break
		}
	}
	//wg.Wait()
}
