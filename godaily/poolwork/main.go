package main

import (
	"fmt"
	"sync"
)

var (
	c1 = make(chan testData)
	gp sync.WaitGroup
)

func main() {
	// TODO:线程数量=数据数量/200+1
	for i := 0; i < (len(list)/200 + 1); i++ {
		go func(i int) {
			for {
				select {
				case data := <-c1:
					process(i, data)
				}
			}
		}(i)
	}
	for _, item := range list {
		c1 <- item
	}
	gp.Wait()
	fmt.Println("=====done=====")
}

func process(i int, data testData) {
	gp.Add(1)
	defer gp.Done()
	// TODO:具体业务逻辑
	fmt.Println("thread ", i, "=", data)

}

type testData struct {
	ID   int
	Name string
	Age  int
}

var list = []testData{
	{1, "aaa", 19},
	{1, "aaa", 19},
	{1, "aaa", 19},
}
