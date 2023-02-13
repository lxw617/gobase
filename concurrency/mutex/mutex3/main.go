package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	var mux sync.Mutex

	for i := 1; i <= 20; i++ {
		go func(i int) {
			mux.Lock()
			fmt.Printf("第%d个go程获取到锁！\n", i)
			time.Sleep(time.Second)
			mux.Unlock()
		}(i)
		time.Sleep(time.Millisecond)
		fmt.Printf("第%d个go程准备就位\n", i)
	}

	// 睡两秒确保上述50个go程均处于阻塞状态
	time.Sleep(time.Second * 2)

	// 监听go程数量
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println("此时go程的数量是:", runtime.NumGoroutine())
			time.Sleep(time.Second)
		}
	}()

	// 每隔五秒创建10个新的go程，循环5次，同样无法模拟新的go程会比之前50个go程优先抢到锁
	for i := 0; i < 5; i++ {
		defer wg.Done()
		for j := 0; j < 10; j++ {
			go func(i int, j int) {
				mux.Lock()
				// fmt.Println("新go程获取到锁！")
				fmt.Printf("新go-%d-%d获取到锁！\n", i, j)
				mux.Unlock()
			}(i, j)
			time.Sleep(time.Millisecond)
			fmt.Printf("新go-%d-%d准备好了\n", i, j)
		}
		time.Sleep(time.Second * 5)
	}
	wg.Wait()
}
