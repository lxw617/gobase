package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	// 并发WaitGroup使用
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(index int) {
			for i := 1; i <= 5; i++ {
				fmt.Printf("第%d次执行,i的值为:%d\n", index, i)
			}
			wg.Done()
		}(i)
	}
	//阻塞,知道WaitGroup队列中所有任务执行结束时自动解除阻塞
	fmt.Println("开始阻塞")
	wg.Wait()
	fmt.Println("任务执行结束,解除阻塞")

}
