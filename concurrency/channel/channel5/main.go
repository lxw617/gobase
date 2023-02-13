package main

import (
	"fmt"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。

func main() {
	chArr := [4]chan struct{}{
		make(chan struct{}),
		make(chan struct{}),
		make(chan struct{}),
		make(chan struct{}),
	}

	for i := 0; i < 4; i++ {
		go func(i int) {
			for {
				<-chArr[i%4]
				fmt.Printf("i am %d\n", i)

				time.Sleep(1 * time.Second)
				chArr[(i+1)%4] <- struct{}{}
			}
		}(i)
	}

	chArr[0] <- struct{}{}
	select {}
}
