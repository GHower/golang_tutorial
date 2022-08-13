package main

import "fmt"

type A int

func main() {
	//// 基本用法
	//a := add(1, 2)
	//
	//// 同类型参数，
	//b := add(1, 2)
	//
	//// 多值返回
	//c, d := swap("str1", "str2")
	//
	//// 返回值命名
	//x, y := split(2)
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Printf("%s---%s", c, d)
	//fmt.Printf("%d---%d", x, y)
	//
	//var a A = 5
	//DoSomething1(&a)
	//DoSomething2(a)

	arr := []int{1, 2, 3}
	dongtaicanshu(arr[1:]...)
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

//func DoSomething1(a *A) {
//b := a
//var aa A = 1
//b = &aa
//println(a)
//println(*a)
//}

//func DoSomething2(a A) {
//	b := &a
//	//var aa A = 2
//	//b = &aa
//	println(a)
//	//println(a)
//}

func dongtaicanshu(arr ...int) {
	for i, v := range arr {
		fmt.Printf("%d:%d\n", i, v)
	}
}
