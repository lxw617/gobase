package main

import (
	"fmt"
	"sync"
	"unsafe"
)

// 查询 WaitGroup 的当前的计数值，包括32位和64位的区分，因为没有32位测试环境 所以仅测试了64位
type WaitGroup struct {
	sync.WaitGroup
}

func (receiver *WaitGroup) GetCounter32() uint32 {
	pointer := unsafe.Pointer(&receiver.WaitGroup)
	return *(*uint32)(unsafe.Pointer(uintptr(pointer) + 8))
}
func (receiver *WaitGroup) GetCounter64() uint32 {
	pointer := unsafe.Pointer(&receiver.WaitGroup)
	return *(*uint32)(unsafe.Pointer(uintptr(pointer) + 4))
}

func main() {
	a := WaitGroup{}
	a.Add(19999)
	fmt.Println(a.GetCounter64())
}
