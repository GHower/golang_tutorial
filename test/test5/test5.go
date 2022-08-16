package main

import (
	"runtime"
	"unsafe"
)

func main() {
	var x int = 43
	var p uintptr = uintptr(unsafe.Pointer(&x))

	runtime.GC()
	println(p)
	var px *int = (*int)(unsafe.Pointer(p))
	println(*px)
}
