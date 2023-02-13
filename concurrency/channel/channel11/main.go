package main

import (
	"fmt"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
const chanNum int = 4

func main() {
	chanArr := make([]chan int, chanNum)
	for i := 0; i < chanNum; i++ {
		ch := make(chan int, 1)
		chanArr[i] = ch
	}

	chanArr[0] <- 1
	for i := 0; i < chanNum; i++ {
		nextChanIdx := (i + 1) % chanNum
		go func(cur, next chan int, idx int) {
			for {
				<-cur
				time.Sleep(1 * time.Second)
				fmt.Printf("%d\n", idx+1)
				next <- 1
			}
		}(chanArr[i], chanArr[nextChanIdx], i)
	}
	time.Sleep(60 * time.Second)
}
