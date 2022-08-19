package main

import "fmt"

func main() {

	go func() {
		for i := 1; i <= 100; i++ {
			fmt.Printf("第%d次执行\n", i)
		}
	}()
	for i := 1; i <= 10; i++ {
		fmt.Println("sqwsd")
	}
}
