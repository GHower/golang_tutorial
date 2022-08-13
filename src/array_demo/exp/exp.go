package main

import "fmt"

func main() {
	//a, b := re2string("abcd", 2)
	//fmt.Printf("%s:%s", a, b)

	//reverseString("Google")
	//oneCharToString([]int{1, 2, 3, 3, 3, 3, 4})

	arr := intMul10(mul10, 1, 2, 3)
	for _, v := range arr {
		fmt.Print(" ", v)
	}
}

// 7.12
func re2string(s string, i int) (string, string) {
	return s[:i], s[i:]
}

// 7.14
func reverseString(s string) {
	var slice = []byte(s)
	for i := 0; i < (len(slice) / 2); i++ {
		slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}
	fmt.Printf("%s", slice)
}

// 7.15
func oneCharToString(arr []int) {
	barr := make([]int, 1)
	barr[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			barr = append(barr, arr[i])
		}
	}
	fmt.Printf("%d", barr)
}

// 7.17
func intMul10(f func(x int) int, arr ...int) []int {
	for i, v := range arr {
		arr[i] = f(v)
	}

	return arr
}
func mul10(a int) int {
	return a * 10
}
