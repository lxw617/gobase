package main

import (
	"fmt"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
func main() {
	ch := make(chan interface{}, 0)
	ch0 := ch
	max := 20
	for i := 0; i < max; i++ {
		i := i
		if i == max-1 {
			go fn(i, ch, ch0)
			continue
		}
		chNext := make(chan interface{}, 0)
		go fn(i, ch, chNext)
		ch = chNext
	}
	ch0 <- 0
	time.Sleep(time.Minute)
}

func fn(i int, ch <-chan interface{}, chNext chan<- interface{}) {
	for {
		if i == 0 {
			time.Sleep(time.Second)
		}
		<-ch
		fmt.Printf("ch:%p, chNext:%p, i:%v \n", ch, chNext, i)
		chNext <- 0
	}
}
