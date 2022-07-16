package main

import "fmt"

func main() {
	// 基本用法
	a := add(1, 2)

	// 同类型参数，
	b := add(1, 2)

	// 多值返回
	c, d := swap("str1", "str2")

	// 返回值命名
	x, y := split(2)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%s---%s", c, d)
	fmt.Printf("%d---%d", x, y)
}

func split(i int) (x, y int) {
	x = i
	y = i * i
	return
}

func swap(x, y string) (string, string) {
	return y, x
}

func add(x int, y int) int {
	return x + y
}
func add2(x, y int) int {
	return x + y
}
