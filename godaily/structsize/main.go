package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	b byte
	i int64
	u uint16
}

func main() {
	var t T
	fmt.Println(unsafe.Sizeof(t))
}
