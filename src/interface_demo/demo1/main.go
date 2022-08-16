package main

import (
	"fmt"
	"reflect"
)

type Reader interface {
	Read(path string) string
}

type A struct {
}

func (a *A) Read(path string) string {
	return path
}
func main() {
	var r Reader
	t1 := reflect.TypeOf(r)
	fmt.Println(t1)
	r = &A{}
	p := r.Read("1.txt")
	t2 := reflect.TypeOf(r)
	fmt.Println(t2)
	println(p)
}
