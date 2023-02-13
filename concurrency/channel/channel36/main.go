package main

import (
	"fmt"
	"time"
)

/*有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，
要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
*/
type Tasker struct {
	Id            int
	SignalCh      chan struct{}
	BroadcastChan chan struct{}
}

func NewTask(id int, ch chan struct{}, bch chan struct{}) *Tasker {
	return &Tasker{
		Id:            id,
		SignalCh:      ch,
		BroadcastChan: bch,
	}
}

func (t *Tasker) say() {
	for {
		<-t.SignalCh
		fmt.Println(t.Id)
		time.Sleep(time.Second * 1)
		t.BroadcastChan <- struct{}{}
	}
}

func ChannelArrange(rNum int) {
	chanArray := make([]chan struct{}, 0, rNum)
	for i := 0; i < rNum; i++ {
		chanArray = append(chanArray, make(chan struct{}))
	}
	worker := make([]*Tasker, 0, rNum)
	for i := 0; i < rNum; i++ {
		worker = append(worker, NewTask(i+1, chanArray[(i-1+rNum)%rNum], chanArray[(i+rNum)%rNum]))
	}

	for i := 0; i < rNum; i++ {
		curIndex := i
		go func() {
			worker[curIndex].say()
		}()
	}
	chanArray[rNum-1] <- struct{}{}

	select {}
}
