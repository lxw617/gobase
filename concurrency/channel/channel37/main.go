package main

import (
	"log"
	"time"
)

/*有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，
要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
*/
func main() {
	var a = make(chan int, 1)
	var b = make(chan int, 1)
	var c = make(chan int, 1)
	var d = make(chan int, 1)
	var e = make(chan string)
	go func() {
		for {
			flag := <-d
			log.Println(1)
			a <- flag
		}
	}()
	go func() {
		for {
			flag := <-a
			log.Println(2)
			b <- flag
		}
	}()
	go func() {
		for {
			flag := <-b
			log.Println(3)
			c <- flag
		}
	}()
	go func() {
		for {
			flag := <-c
			log.Println(4)
			time.Sleep(time.Second)
			d <- flag
		}
	}()
	d <- 1
	<-e
}
