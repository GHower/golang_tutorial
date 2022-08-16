package main

import (
	"errors"
	"fmt"
)

func main() {
	v := Foo()
	fmt.Println(v)
}
func Foo() (err error) {
	if err := errors.New("ss"); err != nil {
		return
	}
	return
}
