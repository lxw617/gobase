package main

import (
	"fmt"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
// 因为channel的等待队列是先入先出的，所以我这边取巧地在goroutine前加一个等待时间，保证1~4的goroutine，他们在同个chan阻塞时是有序的

func main() {
	ch := make(chan struct{})
	for i := 1; i <= 4; i++ {
		go func(index int) {
			time.Sleep(time.Duration(index*10) * time.Millisecond)
			for {
				<-ch
				fmt.Printf("I am No %d Goroutine\n", index)
				time.Sleep(time.Second)
				ch <- struct{}{}
			}
		}(i)
	}
	ch <- struct{}{}
	time.Sleep(time.Minute)
}
