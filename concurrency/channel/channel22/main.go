package main

import (
	"fmt"
	"sync"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
func chanHadle(first bool, num int, send chan bool, recv chan bool) {
	if first {
		time.Sleep(time.Second)
		fmt.Printf("输出%d \n", num)
		send <- true
	}
	for {
		<-recv
		time.Sleep(time.Second)
		fmt.Printf("输出%d \n", num)
		send <- true
	}

}

func main() {
	var chan1 chan bool = make(chan bool)
	var chan2 chan bool = make(chan bool)
	var chan3 chan bool = make(chan bool)
	var chan4 chan bool = make(chan bool)
	var wg sync.WaitGroup

	wg.Add(4)

	go chanHadle(true, 1, chan1, chan4)
	go chanHadle(false, 2, chan2, chan1)
	go chanHadle(false, 3, chan3, chan2)
	go chanHadle(false, 4, chan4, chan3)

	wg.Wait()
}
