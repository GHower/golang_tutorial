package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
)

type Task struct {
	index int
	nums  []int
	sum   int
	wg    *sync.WaitGroup
}

func (t *Task) Do() {
	for _, num := range t.nums {
		t.sum += num
	}

	t.wg.Done()
}

func taskFunc(data interface{}) {
	task := data.(*Task)
	task.Do()
	fmt.Printf("task:%d sum:%d\n", task.index, task.sum)
}
func main() {
	p, _ := ants.NewPoolWithFunc(10, taskFunc)
	defer p.Release()
}
