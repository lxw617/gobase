package main

import (
	"fmt"
	"time"
)

func main() {
	// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
	signChan1 := make(chan struct{})
	signChan2 := make(chan struct{})
	signChan3 := make(chan struct{})
	signChan4 := make(chan struct{})
	mainSignChan := make(chan struct{})

	for i := 1; i <= 4; i++ {
		go func(i int) {
			for {
				select {
				case <-signChan1:
					fmt.Println(1)
					time.Sleep(1 * time.Second)
					signChan2 <- struct{}{}
				case <-signChan2:
					fmt.Println(2)
					time.Sleep(1 * time.Second)
					signChan3 <- struct{}{}
				case <-signChan3:
					fmt.Println(3)
					time.Sleep(1 * time.Second)
					signChan4 <- struct{}{}
				case <-signChan4:
					fmt.Println(4)
					time.Sleep(1 * time.Second)
					signChan1 <- struct{}{}
				}
			}
		}(i)
	}
	signChan1 <- struct{}{}
	<-mainSignChan
}
