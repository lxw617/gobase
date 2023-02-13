package main

import (
	"fmt"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
func main() {
	chArr := []chan struct{}{
		make(chan struct{}),
		make(chan struct{}),
		make(chan struct{}),
		make(chan struct{}),
		make(chan struct{}),
		make(chan struct{}),
	}
	for k, _ := range chArr {
		if k == len(chArr)-1 {
			go goon(chArr[k], chArr[0], k+1)
		} else {
			go goon(chArr[k], chArr[k+1], k+1)
		}
	}

	chArr[0] <- struct{}{}
	select {}

}

func goon(ch chan struct{}, ch2 chan struct{}, index int) {
	time.Sleep(time.Duration(index*10) * time.Millisecond)
	for {
		<-ch
		fmt.Printf("I am No %d Goroutine\n", index)
		time.Sleep(time.Second)
		ch2 <- struct{}{}
	}
}
