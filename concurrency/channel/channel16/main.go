package main

import (
	"fmt"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
func main() {
	num := make(chan int, 4)
	go func() {
		for {
			for i := 1; i <= 4; i++ {
				num <- i
				time.Sleep(time.Second * 1)
			}
		}
	}()

	go func() {
		for {
			fmt.Println(<-num)
		}
	}()

	select {}
}
