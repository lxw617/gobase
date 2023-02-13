package remetux

// import (
// 	"fmt"
// 	"sync"
// 	"sync/atomic"
// )

// const (
// 	READ  = 0
// 	WRITE = 1
// )

// type RWMutex struct {
// 	sync.RWMutex
// }

// func (rw *RWMutex) TryLock(mode int32) bool {

// 	n := atomic.LoadInt32(&rw.readerCount)
// 	if n < 0 {
// 		return false
// 	}
// 	if mode == 0 && atomic.CompareAndSwapInt32(n, n+1) == n {
// 		return true
// 	} else if mode == 1 && atomic.CompareAndSwapInt32(0, -rwmutexMaxReaders) == 0 {
// 		rw.Lock()
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func (rw *RWMutex) ExistWriter() bool {
// 	return atomic.LoadInt32(&rw.readerCount) < 0
// }
// func (rw *RWMutex) GetReaderNum() int32 {
// 	n := atomic.LoadInt32(&rw.readerCount)
// 	if n < 0 {
// 		return atomic.LoadInt32(&rw.readerWait)
// 	} else {
// 		return n
// 	}
// }
// func main() {
// 	a := RWMutex{}

// 	fmt.Println(a.GetReaderNum())
// 	fmt.Println(a.ExistWriter())
// }
