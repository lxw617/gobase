package main

import (
	"fmt"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
type Worker struct{}

// 把channel1的例子重新用另外一种方式来写，需求：打印1，2，3，4 周而复始

func startWorker(workerId int, startWorkChan, nextWorkChan chan Worker) {
	for {
		// 工人接到了工作
		needDoWork := <-startWorkChan
		fmt.Println("打印--->", workerId+1)
		time.Sleep(time.Second)
		// 当前工人累了，把工作交给下一个工人
		nextWorkChan <- needDoWork
	}
}

func main() {
	chs := []chan Worker{make(chan Worker), make(chan Worker), make(chan Worker), make(chan Worker)}
	// 创建四个工人
	var myWorkerNum = 4

	// 给四个工人分配工作
	for i := 0; i < myWorkerNum; i++ {
		go startWorker(i, chs[i], chs[(i+1)%myWorkerNum])
	}

	// 首先给第一个工人工作去做事
	chs[0] <- struct{}{}
	select {}
}
