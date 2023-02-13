package main

import (
	"fmt"
	"sync"
	"time"
)

/*有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，
要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
*/
func main() {
	const chanNum int = 4
	chanArry := make([]chan int, chanNum)
	for i := 0; i < chanNum; i++ {
		chanArry[i] = make(chan int, 1)
	}

	elapseSec := 10

	wg := new(sync.WaitGroup)
	wg.Add(chanNum)

	quitCh := make(chan struct{})
	chanArry[0] <- 1
	for i := 0; i < chanNum; i++ {
		nextIdx := (i + 1) % chanNum
		go func(curCh, nextCh chan int, idx int, quitCh chan struct{}, wg *sync.WaitGroup) {
		Loop:
			for {
				select {
				case val := <-curCh:
					fmt.Printf("I am goroutine:%d,val:%d\n", idx, val)
					time.Sleep(time.Second)
					nextCh <- (val + 1)
				case <-quitCh:
					fmt.Printf("-->goroutine:%d exit!\n", idx)
					break Loop
				}
			}
			wg.Done()
		}(chanArry[i], chanArry[nextIdx], i+1, quitCh, wg)
	}

	select {
	case <-time.After(time.Second * time.Duration(elapseSec)):
		fmt.Println("-->begin close goroutine....")
		close(quitCh)
	}
	wg.Wait()
	fmt.Println("all goroutine exit!,will be exit!")
}
