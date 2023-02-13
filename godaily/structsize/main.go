package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type T struct {
		b byte

		i int64
		u uint16
	}
	var t T
	fmt.Println(unsafe.Sizeof(t))
}
