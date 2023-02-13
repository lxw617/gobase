package main

import (
	"fmt"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})
	ch4 := make(chan struct{})

	go func() {
		for {
			<-ch1
			time.Sleep(time.Second)
			fmt.Println("No.1")
			ch2 <- struct{}{}
		}
	}()

	go func() {
		for {
			<-ch2
			time.Sleep(time.Second)
			fmt.Println("No.2")
			ch3 <- struct{}{}
		}
	}()

	go func() {
		for {
			<-ch3
			time.Sleep(time.Second)
			fmt.Println("No.3")
			ch4 <- struct{}{}
		}
	}()

	go func() {
		for {
			<-ch4
			time.Sleep(time.Second)
			fmt.Println("No.4")
			ch1 <- struct{}{}
		}
	}()

	ch1 <- struct{}{}
	select {}
}
