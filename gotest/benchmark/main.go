package benchmark

import "fmt"

func TestB() {
	go func() {
		for i := 1; i <= 100; i++ {
			fmt.Printf("第%d次执行", i)
		}
	}()
}
