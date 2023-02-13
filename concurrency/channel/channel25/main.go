package main

import (
	"fmt"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
func main() {
	chans := [4]chan struct{}{}
	for i := range chans {
		chans[i] = make(chan struct{})
	}
	task := func(id int) {
		c := chans[id-1]
		nextC := chans[id%len(chans)]
		for {
			<-c
			fmt.Println(id)
			time.Sleep(time.Second)
			nextC <- struct{}{}
		}
	}

	for i := 1; i <= 4; i++ {
		go task(i)
	}
	chans[0] <- struct{}{}
	time.Sleep(100 * time.Second)
}
