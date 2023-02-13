package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
func main() {

	wg := sync.WaitGroup{}

	chans := []chan int{make(chan int, 1), make(chan int, 1), make(chan int, 1), make(chan int, 1)}

	n := len(chans)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	proc := func(i int) {
		j := i + 1
		defer wg.Done()
		for {
			select {
			case <-chans[i%n]:
				fmt.Printf("%d -> ", j)
				time.After(time.Second)

				if ctx.Err() != nil {
					return
				}

				chans[j%n] <- j % n
			case <-ctx.Done():
				return
			}
		}
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go proc(i)
	}

	chans[0] <- 0

	wg.Wait()
}
