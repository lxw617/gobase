package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	mu.Lock()
	foo(mu)
	mu.Unlock()
	fmt.Println("main...")
}
func foo(mu sync.Mutex) {
	mu.Lock()
	fmt.Println("foo...")
	mu.Unlock()
}
