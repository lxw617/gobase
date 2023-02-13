package main

import (
	"fmt"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)

	go func() {
		for {
			<-ch1
			fmt.Println(1)
			ch2 <- 1
		}
	}()
	go func() {
		for {
			<-ch2
			fmt.Println(2)
			ch3 <- 1
		}
	}()
	go func() {
		for {
			<-ch3
			fmt.Println(3)
			ch4 <- 1
		}
	}()
	go func() {
		for {
			<-ch4
			fmt.Println(4)
		}
	}()

	for {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}

}
