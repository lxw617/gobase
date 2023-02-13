// 打印任意数

package main

import (
	"fmt"
	"time"
)

const (
	routineNum = 7
)

func main() {
	var single chan struct{}
	var channelList = make([]chan struct{}, routineNum)
	for i := 0; i < routineNum; i++ {
		channelList[i] = make(chan struct{})
	}
	for i := 0; i < routineNum; i++ {
		go func(i int) {
			for {
				<-channelList[i%routineNum]
				fmt.Println((i % routineNum) + 1)
				time.Sleep(1 * time.Second)
				channelList[(i+1)%routineNum] <- struct{}{}
			}
		}(i)
	}
	channelList[0] <- struct{}{}
	<-single
}
