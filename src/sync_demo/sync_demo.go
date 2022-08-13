package main

import (
	"sync"
)

func main() {
	info := Info{str: "22"}
	demo1(&info)
}

type Info struct {
	mu  sync.Mutex
	str string
}

func demo1(info *Info) {
	// 上锁
	info.mu.Lock()
	info.str = "newstr"
	// 解锁
	info.mu.Unlock()
}
