package main_test

import (
	"fmt"
	"os"
	"testing"
)

var temp = make(map[string][]string)

func Test222(t *testing.T) {

	file, err := os.OpenFile("./", os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	files, err := file.Readdir(-1)
	if err != nil {
		return
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
