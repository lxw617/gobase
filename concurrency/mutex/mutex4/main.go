package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 设置一个共享区
	number := 0

	var mux sync.Mutex

	// var wg sync.WaitGroup

	// wg.Add(20)
	for i := 1; i <= 100; i++ {
		go func(i int) {
			// defer wg.Done()
			mux.Lock()
			number++
			fmt.Printf("第%d个go程把%d加一为%d\n", i, number-1, number)
			mux.Unlock()
			time.Sleep(time.Second * 2)
		}(i)
	}
	// wg.Wait()

	// 睡两秒确保上述50个go程均处于阻塞状态
	time.Sleep(time.Second * 2)

	// 监听go程数量
	for i := 0; i < 100; i++ {
		go func() {
			mux.Lock()
			number++
			fmt.Printf("把%d加一为%d\n", number-1, number)
			mux.Unlock()
		}()
	}
}
