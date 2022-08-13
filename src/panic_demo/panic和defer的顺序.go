package main

import "fmt"

/**
当panic时，会先执行defer再向上panic，在defer中可以用
recover恢复
*/
func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("packing %s \r\n", e)
		}
	}()
	bad()
	fmt.Println("after")
}
func bad() {
	panic("call bad")
}
func main() {
	println("call")
	test()
	println("end")
}
