package main

import (
	"fmt"
	"testing"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
func Test_orderPrint(t *testing.T) {
	c := make(chan int)
	go func() {
		c <- 1
	}()

	go orderPrint(c, 1)
	go orderPrint(c, 2)
	go orderPrint(c, 3)
	go orderPrint(c, 4)

	select {}
}

func orderPrint(c chan int, i int) {
	for {
		time.Sleep(time.Second)
		cur := <-c
		if cur == i {
			fmt.Println(i)
			if i%4 == 0 {
				c <- 1
			} else {
				c <- i + 1
			}
		} else {
			c <- cur
		}
	}
}
