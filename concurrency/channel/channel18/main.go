package main

import (
	"sync"
	"testing"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
// 并不是每个goroutine打印自己的id
func TestChannel1Practice(t *testing.T) {
	var ch = make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(4)

	go func() {
		ch <- struct{}{}
	}()

	for thread := 1; thread <= 4; thread++ {
		go func(thead int) {
			_, ok := <-ch
			if ok {
				for i := 1; i <= 4; i++ {
					t.Logf("%d: %d", thead, i)
					time.Sleep(1 * time.Second)
				}
				wg.Done()
				ch <- struct{}{}
			}
		}(thread)
	}

	wg.Wait()
	t.Log("finished")
}
