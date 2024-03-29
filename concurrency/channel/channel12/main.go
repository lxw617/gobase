package main

import (
	"fmt"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
func printChan(c chan int) {
	st := <-c
	fmt.Println(st%4 + 1)
	time.Sleep(1 * time.Second)
	c <- st + 1
	go printChan(c)
}

func main() {
	ch := make(chan int, 4)
	ch <- 0
	printChan(ch)

	select {}

}
