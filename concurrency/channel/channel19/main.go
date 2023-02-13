package main

import (
	"context"
	"fmt"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
func tt(ctx context.Context, c1, c2 *chan int) {
	for {
		select {
		case n := <-*c1:
			fmt.Println(n)
			nn := n + 1
			if n == 4 {
				nn = 1
			}
			*c2 <- nn
			//fmt.Printf("c1:%p,c2:%p\n", c1, c2)
		case <-ctx.Done():
			return

		}
	}
}
func PrintInfo() {
	ctx, cancel := context.WithCancel(context.Background())
	c1, c2, c3, c4 := make(chan int, 2), make(chan int, 2), make(chan int, 2), make(chan int, 2)
	fmt.Printf("c1:%p,c2:%p,c3:%p,c4:%p\n", &c1, &c2, &c3, &c4)
	go tt(ctx, &c1, &c2)
	go tt(ctx, &c2, &c3)
	go tt(ctx, &c3, &c4)
	go tt(ctx, &c4, &c1)
	c1 <- 1

	fmt.Println("Hello, 世界")
	time.Sleep(time.Millisecond * 70)
	cancel()

	fmt.Println("Hello, 世界")
}
