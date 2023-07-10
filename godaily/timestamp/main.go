package main

import (
	"fmt"
	"time"
)

func main() {
	s1, _ := time.Parse("2006-01-02T15:04:05+08:00", "0001-01-01T00:00:00+08:05")
	fmt.Println(s1)
	fmt.Println(s1.Unix())

	s5, _ := time.Parse("2006-01-02T15:04:05+08:00", "2023-07-10T14:43:35+08:00")
	fmt.Println(s5)
	fmt.Println(s5.Unix())

	s2 := time.Unix(1688955651, 0)
	fmt.Println(s2)
	fmt.Println(s2.Unix())

	s3 := time.Date(2023, time.Month(7), 10, 2, 30, 30, 0, time.Local)
	fmt.Println(s3)
	fmt.Println(s3.Unix())

	s4 := time.Date(0001, time.Month(1), 1, 0, 0, 0, 0, time.Local)
	fmt.Println(s4)
	fmt.Println(s4.Unix())
}
