package main

import "fmt"

func main() {
	var arr1 = [5]int{1, 2, 3}
	var arr2 = new([5]int)
	arr2[3] = 100
	arr1 = *arr2
	arr1[2] = 100
	fmt.Println(arr1, arr2)
}
