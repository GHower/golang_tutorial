package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"os"
	"sync"
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	const (
		DataSize    = 10000
		DataPerTask = 100
	)

	nums := make([]int, DataSize, DataSize)
	//for i := range nums {
	//	nums[i] = rand.Intn(1000)
	//}
	fmt.Println(nums)
}

// 设置最大等待队列长度
func Test_main2(t *testing.T) {
	p, _ := ants.NewPool(4, ants.WithMaxBlockingTasks(2))
	defer p.Release()

	var wg sync.WaitGroup
	wg.Add(8)
	for i := 1; i <= 8; i++ {
		go func(i int) {
			err := p.Submit(wrapper(i, &wg))
			if err != nil {
				fmt.Printf("task:%d err:%v\n", i, err)
				wg.Done()
			}
		}(i)
	}

	wg.Wait()
}

//非阻塞
func Test_main3(t *testing.T) {
	p, _ := ants.NewPool(2, ants.WithNonblocking(true))
	defer p.Release()

	var wg sync.WaitGroup
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		err := p.Submit(wrapper(i, &wg))
		if err != nil {
			fmt.Printf("task:%d err:%v\n", i, err)
			wg.Done()
		}
	}

	wg.Wait()
}

// 默认panic处理器
func Test_main4(t *testing.T) {
	p, _ := ants.NewPool(2)
	defer p.Release()

	var wg sync.WaitGroup
	wg.Add(3)
	for i := 1; i <= 2; i++ {
		p.Submit(wrapper2(i, &wg))
	}

	time.Sleep(1 * time.Second)
	p.Submit(wrapper(3, &wg))
	p.Submit(wrapper(5, &wg))
	wg.Wait()
}

// 自定义panic处理器
func Test_main4_2(t *testing.T) {
	panicH := func(err interface{}) {
		fmt.Fprintln(os.Stderr, err)
	}

	p, _ := ants.NewPool(2, ants.WithPanicHandler(panicH))
	defer p.Release()

	var wg sync.WaitGroup
	wg.Add(3)
	for i := 1; i <= 2; i++ {
		p.Submit(wrapper2(i, &wg))
	}

	time.Sleep(1 * time.Second)
	p.Submit(wrapper(3, &wg))
	p.Submit(wrapper(5, &wg))
	wg.Wait()
}
