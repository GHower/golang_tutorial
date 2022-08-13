package main

import "fmt"

func main() {
	//sl_from := []int{1, 2, 3, 4}
	//sl := sl_from[1:]
	//sl2 := append(sl, 5,5,5)
	//fmt.Println(sl_from)
	//fmt.Println(sl)
	//fmt.Println(sl2)
	//
	//fmt.Println()
	//sl[1] = 9
	//sl2[1] = 10
	//fmt.Println(sl_from)
	//fmt.Println(sl)
	//fmt.Println(sl2)

	sli := make([]int, 0, 4)
	sli2 := append(sli, 1, 2, 3)
	fmt.Println(sli)
	fmt.Println(sli2)
}
