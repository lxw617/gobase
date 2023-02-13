package main

import (
	"fmt"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。

func f(i int, input <-chan int, output chan<- int) {
	for {
		<-input
		fmt.Println(i)
		time.Sleep(time.Second)
		output <- 1
	}
}
func main() {
	c := [4]chan int{}
	for i := range []int{1, 2, 3, 4} {
		c[i] = make(chan int)
	}
	go f(1, c[3], c[0])
	go f(2, c[0], c[1])
	go f(3, c[1], c[2])
	go f(4, c[2], c[3])
	c[3] <- 1
	select {}
}

// func TestChannelPlan(t *testing.T) {
// 	c := [4]chan int{}
// 	for i := range []int{1, 2, 3, 4} {
// 		c[i] = make(chan int)
// 	}
// 	go f(1, c[3], c[0])
// 	go f(2, c[0], c[1])
// 	go f(3, c[1], c[2])
// 	go f(4, c[2], c[3])
// 	c[3] <- 1
// 	select {}
// }
