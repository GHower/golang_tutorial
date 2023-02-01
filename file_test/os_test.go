package file_test

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func Test_os_create(t *testing.T) {
	filepath := "D:/goland_projects/golang_tutorial/file_test/test/1.txt"
	split, file := path.Split(filepath)
	fmt.Println(split, file)

	dir := path.Dir(filepath)
	fmt.Println(dir)

	err := os.MkdirAll(dir, 0611)
	if err != nil {
		return
	}
	// 若目录不存在会报错
	_, err = os.Create(filepath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func Test_os_mkdirall(t *testing.T) {
	filepath := "D:/goland_projects/golang_tutorial/file_test/test2/2.txt"
	// mkdirAll 尽可能创建目录，多次执行不会报错
	err := os.MkdirAll(filepath, 0611)
	if err != nil {
		return
	}
}
