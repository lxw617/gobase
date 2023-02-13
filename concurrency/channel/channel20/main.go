package main

import (
	"fmt"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
const n = 4

func main() {
	var chanArray = [n]chan int{}
	for i := 0; i < n; i++ {
		chanArray[i] = make(chan int)
	}
	for i := 0; i < n; i++ {
		go func(ix int) {
			for v := range chanArray[ix] {
				fmt.Println(v + 1)
			}
		}(i)
	}

	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Second)
			chanArray[i%n] <- i % n
		}
	}()

	for {
	}
}
