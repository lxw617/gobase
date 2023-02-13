package main

import (
	"fmt"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
func main() {
	// 4个goroutine
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})
	ch4 := make(chan struct{})

	go func() {
		for {
			<-ch1
			fmt.Println("1")
			time.Sleep(time.Second)
			ch2 <- struct{}{}
			// <-ch1
		}
	}()

	go func() {
		for {
			<-ch2
			fmt.Println("2")
			time.Sleep(time.Second)
			ch3 <- struct{}{}
		}
	}()

	go func() {
		for {
			<-ch3
			fmt.Println("3")
			time.Sleep(time.Second)
			ch4 <- struct{}{}
		}
	}()

	go func() {
		for {
			<-ch4
			fmt.Println("4")
			time.Sleep(time.Second)
			ch1 <- struct{}{}
		}
	}()

	ch1 <- struct{}{}
	select {}
}
