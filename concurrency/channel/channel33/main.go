package main

import (
	"fmt"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
func Print(recv chan int) {
	for {
		index, ok := <-recv
		if !ok {
			return
		}
		fmt.Println(index)
		index++
		if index == 5 {
			index = 1
		}

		time.Sleep(1 * time.Second)
		recv <- index
	}
}

func main() {
	c := make(chan int)
	go Print(c)
	go Print(c)
	go Print(c)
	go Print(c)
	c <- 1
	time.Sleep(20 * time.Second)
	close(c)
}
