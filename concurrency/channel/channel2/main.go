package main

import (
	"fmt"
	"time"
)

func main() {
	// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)
	go func() {
		for {
			fmt.Println("I'm goroutine 1")
			time.Sleep(1 * time.Second)
			ch2 <- 1 //I'm done, you turn
			<-ch1
		}
	}()

	go func() {
		for {
			<-ch2
			fmt.Println("I'm goroutine 2")
			// time.Sleep(1 * time.Second)
			ch3 <- 1
		}

	}()

	go func() {
		for {
			<-ch3
			fmt.Println("I'm goroutine 3")
			// time.Sleep(1 * time.Second)
			ch4 <- 1
		}

	}()

	go func() {
		for {
			<-ch4
			fmt.Println("I'm goroutine 4")
			// time.Sleep(1 * time.Second)
			ch1 <- 1
		}

	}()

	select {}
}
