package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	demo0()
}

// sizeof
func d1() {
	var p float64 = 99
	// uintptr
	fmt.Println(reflect.TypeOf(unsafe.Sizeof(p)))
	// 8
	fmt.Println(unsafe.Sizeof(p))
}

// sizeof传入指针类型的对象
type W struct {
	a byte
	b int32
	c int64
}

func d2() {
	var w *W
	fmt.Println(unsafe.Sizeof(*w))
}

// 对齐Alignof
func alignofDemo() {
	var w *W
	a := unsafe.Alignof(w.a) // type byte
	b := unsafe.Alignof(w.b) // type int32
	c := unsafe.Alignof(w.c) // type int64
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}

// 综合示例
type T struct {
	t1 byte
	t2 int32
	t3 int64
	t4 string
	t5 bool
}

func demo0() {
	fmt.Println("----------unsafe.Pointer---------")
	t := &T{1, 2, 3, "this is a example", true}
	ptr := unsafe.Pointer(t)
	t1 := (*byte)(ptr)
	fmt.Println(*t1)
	t2 := (*int32)(unsafe.Pointer(uintptr(ptr) + unsafe.Offsetof(t.t2)))
	*t2 = 99
	fmt.Println(t)
	t3 := (*int64)(unsafe.Pointer(uintptr(ptr) + unsafe.Offsetof(t.t3)))
	fmt.Println(*t3)
	*t3 = 123
	fmt.Println(t)
}
