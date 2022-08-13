package main

import (
	"fmt"
	"reflect"
)

// 反射案例1
func reflectTest01(b interface{}) {
	rtpe := reflect.TypeOf(b)

	v := reflect.ValueOf(b)
	// 注意rtpe不是字符串类型，v也不是int类型的值,不能参与运算
	fmt.Println(rtpe, v)
}

func main() {
	// 通过反射对[基本类型，interface{},reflect.value]进行操作
	var num int = 100
	reflectTest01(num)
}
