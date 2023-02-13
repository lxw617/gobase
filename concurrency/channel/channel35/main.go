package main

import (
	"fmt"
	"testing"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
func TestFourGo(t *testing.T) {
	count := 6
	ch := make([]chan bool, 0)
	for i := 0; i < count; i++ {
		ch = append(ch, make(chan bool))
	}

	gor := func(ch1 chan bool, ch2 chan bool, num int8) {
		for {
			<-ch1
			time.Sleep(time.Second)
			fmt.Println(num)
			ch2 <- true
		}
	}
	for i := 0; i < count; i++ {
		n := i + 1
		if i == count-1 {
			n = 0
		}
		go gor(ch[i], ch[n], int8(i+1))
	}
	ch[0] <- true

	time.Sleep(time.Hour)
}
