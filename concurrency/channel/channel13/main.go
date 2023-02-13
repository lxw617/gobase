package main

import (
	"fmt"
	"time"
)

const N int = 4

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
func main() {
	var chans [N]chan struct{}
	for i := 0; i < N; i++ {
		chans[i] = make(chan struct{})
	}

	for i := 0; i < N; i++ {
		go func(i int) {
			for {
				<-chans[i]
				fmt.Print(i + 1)
				time.Sleep(time.Second)
				chans[(i+1)%N] <- struct{}{}
			}
		}(i)
	}

	chans[0] <- struct{}{}

	select {}
}
