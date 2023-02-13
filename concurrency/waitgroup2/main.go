package main

import (
	"log"
	"sync"
	"time"
	"unsafe"
)

func main() {
	// 获取 WaitGroup 的计数值
	var wg sync.WaitGroup

	wg.Add(20)

	// state1 [3]uint32，当前是 64bit 的
	// state1[0]：Waiter数目，也就是调用了 Wait() 的 goroutine 的数量
	// state1[1]：计数值
	for i := 10; i > 0; i-- {
		go func(i int) {
			wg.Wait()
		}(i)
	}

	time.Sleep(1 * time.Second)
	ptr := (*uint64)(unsafe.Pointer((uintptr(unsafe.Pointer(&wg)))))
	counter := int32(*ptr >> 32)
	waiters := uint32(*ptr)

	log.Printf("waiters:%d, counter:%d", waiters, counter)
	wg.Add(-20)
	wg.Wait()
}
