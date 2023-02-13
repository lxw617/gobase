package main

import (
	"fmt"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
func main() {
	var chs [4]chan bool
	for i := 0; i < 4; i++ {
		chs[i] = make(chan bool)
		go printID(i, chs[i])
	}

	c := 0
	for {
		id := c % 4
		chs[id] <- true
		time.Sleep(time.Second)
		c++
	}
}

func printID(id int, ch chan bool) {
	for {
		select {
		case <-ch:
			fmt.Println(id)
		}
	}
}
