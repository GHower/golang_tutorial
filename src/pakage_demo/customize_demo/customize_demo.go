package main

import (
	"fmt"
	"golang_tutorial/src/pakage_demo/customize_demo/pack1"
)
import p "golang_tutorial/src/pakage_demo/customize_demo/pack1"

func main() {
	var test1 string
	test1 = pack1.ReturnStr()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("Integer from package1: %d\n", pack1.Pack1Int)
	// fmt.Printf("Float from package1: %f\n", pack1.pack1Float)

	test2 := p.ReturnStr()
	fmt.Printf("ReturnStr from package1: %s\n", test2)

}
