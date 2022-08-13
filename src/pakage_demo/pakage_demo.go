package main

import (
	"container/list"
	"fmt"
	"unsafe"
)

func main() {
	unsafeDemo()
}

// 练习9.1
//使用 container/list 包实现一个双向链表，将 101、102 和 103 放入其中并打印出来。
func linkedList2() {
	a := list.List{}
	a.PushBack(101)
	a.PushBack(102)
	a.PushBack(103)

	for e := a.Front(); e != nil; e = e.Next() {
		fmt.Print(" ", e.Value)
	}
}

// 练习9.2
// 通过使用 unsafe 包中的方法来测试你电脑上一个整型变量占用多少个字节。
func unsafeDemo() {
	var v int32
	a := unsafe.Sizeof(v)
	fmt.Println(a)
}
