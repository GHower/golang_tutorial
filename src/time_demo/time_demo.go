package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%04d.%02d.%2d\n", t.Year(), t.Month(), t.Day())

	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Printf("%04d.%02d.%2d\n", t.Year(), t.Month(), t.Day())
	// 计算时间, 必须是纳秒计算
	week = 60 * 60 * 24 * 7 * 1e9
	weekFrom := t.Add(week)
	fmt.Println(weekFrom)

	// 格式化时间
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2006 15:04"))
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
}
