package main

import (
	"fmt"
	"math"
)

func main() {
	x := pow(2, 4, 32)
	fmt.Println(x)
}

/**
1. 对比其他高级语言，去除了小括号。
2. 可以在条件句v < lim之前执行简单的语句v ：= math.Pow(x,n)
*/
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
