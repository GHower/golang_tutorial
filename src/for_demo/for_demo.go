package main

import "fmt"

func main() {
	// 基本使用
	func_for_1()
	// 省略内容
	func_for_2()

}

func func_for_2() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}

func func_for_1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
