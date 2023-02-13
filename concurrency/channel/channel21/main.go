package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
func helper(signal chan int, i int) {
	for {
		signal <- i
	}
}
func main() {
	N := 4
	signals := make([]chan int, N)
	for i := range signals {
		signals[i] = make(chan int)
	}
	sw := sync.WaitGroup{}
	defer sw.Wait()
	defer os.Exit(0)
	for i := 0; i < N; i++ {
		sw.Add(1)
		go func(i int) {
			defer sw.Done()
			helper(signals[i], i)
		}(i)

	}
	for {
		for i := 0; i < N; i++ {
			fmt.Println(<-signals[i] + 1)
			time.Sleep(1 * time.Second)
		}
	}

}
