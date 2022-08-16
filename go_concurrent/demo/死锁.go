package test

import (
	"sync"
	"sync/atomic"
)

func main() {

	var wg sync.WaitGroup

	ans := int64(0)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go newGoRoutine(wg, &ans)
	}
	wg.Wait()
}

// 这样写会死锁，是因为wg是值传递，newGoRoutine中的wg.Done并不是main 中的
// 最终导致main中的wg没有执行过done而死锁
func newGoRoutine(wg sync.WaitGroup, i *int64) {
	defer wg.Done()
	atomic.AddInt64(i, 1)
	return
}
