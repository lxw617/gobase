package main

import (
	"fmt"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
/*
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFINGEMENT. IN NO EVENT SHALL THEq
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */

type NumberChan struct {
	Ch            chan int
	ChannelNumber int
}

func (nch *NumberChan) SendNotify() {
	go func() {
		nch.Ch <- nch.ChannelNumber
	}()
}

func (nch *NumberChan) PrintInfo() {
	fmt.Println(nch.ChannelNumber)
	time.Sleep(time.Second)
}

func NewNumberChan(seq int) *NumberChan {
	nch := NumberChan{
		Ch:            make(chan int),
		ChannelNumber: seq,
	}
	return &nch
}

func main() {
	var (
		nch1 = NewNumberChan(1)
		nch2 = NewNumberChan(2)
		nch3 = NewNumberChan(3)
		nch4 = NewNumberChan(4)
	)
	go func() {
		nch1.SendNotify()
	}()
	for {
		select {
		case <-nch1.Ch:
			nch1.PrintInfo()
			nch2.SendNotify()
		case <-nch2.Ch:
			nch2.PrintInfo()
			nch3.SendNotify()
		case <-nch3.Ch:
			nch3.PrintInfo()
			nch4.SendNotify()
		case <-nch4.Ch:
			nch4.PrintInfo()
			nch1.SendNotify()
		}
	}

}
