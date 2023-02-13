package main

import (
	"fmt"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。

func chgoroutine(in, out, stop chan struct{}, n int) {
	for {
		select {
		case <-in:
			fmt.Println(n)
			time.Sleep(time.Second)
			out <- struct{}{}
		case <-stop:
			return
		}
	}
}

func main() {
	ch1 := make(chan struct{}, 0)
	ch2 := make(chan struct{}, 0)
	ch3 := make(chan struct{}, 0)
	ch4 := make(chan struct{}, 0)
	stop := make(chan struct{}, 0)

	go chgoroutine(ch1, ch2, stop, 1)
	go chgoroutine(ch2, ch3, stop, 2)
	go chgoroutine(ch3, ch4, stop, 3)
	go chgoroutine(ch4, ch1, stop, 4)

	ch1 <- struct{}{}

	time.Sleep(time.Second * 20)

	stop <- struct{}{}
}
