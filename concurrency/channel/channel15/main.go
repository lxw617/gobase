package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
func main() {
	const total int64 = 4
	ch := make(chan int64)
	var i, flag int64

	for ; i < total; i++ {
		go func(i int64) {
			for {
				v := <-ch
				if v%total == i {
					fmt.Println(i+1, v)
					time.Sleep(time.Second)
					atomic.AddInt64(&flag, 1)
				}
				ch <- atomic.LoadInt64(&flag)
			}
		}(i)
	}

	ch <- flag
	select {}
}
