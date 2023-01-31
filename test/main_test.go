package main

import (
	"fmt"
	"testing"
)

type N struct {
	name string
}

func Test_test1(t *testing.T) {
	n := &N{name: "ss"}
	m := make(map[string]*N)
	m["n"] = n
	k := getMap(m)
	println(k)
}
func getMap(m interface{}) int {
	mm := m.(map[string]interface{})
	fmt.Println(mm)
	return len(mm)
}
