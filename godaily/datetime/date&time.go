package main

import (
	"fmt"
	"time"
)

func main() {
	// 1．获取当前时间
	now := time.Now()
	fmt.Printf("now=%v now type=%T\n", now, now)
	// 2.通过now可以获取到年月日,时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%vin", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%vln", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	// 格式化日期时间
	// 使用 Printf 或者 SPrintf
	fmt.Printf("当前年月日%d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	datestr := fmt.Sprintf("当前年月日%d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("datestr=%v\n", datestr)
	// 使用 time.Format() 方法完成:
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))
	// "2006/01/02 15:04:05" 这个字符串的各个数字是固定的，必须是这样写。
	// "2006/01/02 15:04:05" 这个字符串各个数字可以自由的组合，这样可以按程序需求来返回时间和日期

	/*
	   时间的常量:在程序中可用于获取指定时间单位的时间，比如想得到100 毫秒100 * time. Millisecond
	   const (
	   Nanosecond Duration = 1 //纳秒
	   Microsecond = 1000 * Nanosecond //微秒
	   Millisecond = 1000 * Microsecond //毫秒
	   Second = 1000 * Millisecond //秒
	   Minute = 60 * Second //分钟
	   Hour = 60 * Minute //小时
	   )
	*/
	// time 的 Unix 和 UnixNano 的方法
	fmt.Printf("unix时间戳=%v unixnano时间戳=%v\n", now.Unix(), now.UnixNano()) // unix时间戳=1675685905 unixnano时间戳=1675685905428000000
	fmt.Println("----------------------------------------------------------------")
	if time.Now().After(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)) && time.Now().Before(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 6, 0, 0, 0, time.Local)) {
		fmt.Println("-----222222222------------------------")
		fmt.Println(time.Now().After(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)))
		fmt.Println(time.Now().Before(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 6, 0, 0, 0, time.Local)))
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local).Format("2006-01-02 15:04:05"))
		fmt.Println(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 6, 0, 0, 0, time.Local).Format("2006-01-02 15:04:05"))

	}
	fmt.Println("-----111111111111------------------")
	fmt.Println(time.Unix(1684225865, 0).Format("2006-01-02 15:04:05"))
	a := fmt.Sprintf("%.2f", float64(0)/100)
	fmt.Println(a)
}
