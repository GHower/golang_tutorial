package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	slice1 := a[:int64(math.Round(float64(len(a))/2))]
	for i, v := range slice1 {
		fmt.Println(i, v)
	}
}
