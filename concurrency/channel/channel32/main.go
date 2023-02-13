package main

import (
	"fmt"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
func main() {
	size := 4
	signChans := make([]chan struct{}, size)
	for i := range signChans {
		signChans[i] = make(chan struct{})
	}

	for i := 1; i <= size; i++ {
		go func(id int, sign chan struct{}) {
			for {
				<-sign
				fmt.Println(id)
			}
		}(i, signChans[i-1])
	}

	// controller
	go func() {
		ticker := time.NewTicker(time.Second)
		id := 1
		for {
			<-ticker.C
			signChans[id-1] <- struct{}{}
			id++
			if id > size {
				id = 1
			}
		}
	}()

	select {}
}
