package main

import (
	"fmt"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
type NumChan struct {
	jobs []*Job
}

func (n *NumChan) JobNum(m int) {
	for i := 1; i <= m; i++ {
		job := &Job{
			ID:   i,
			Jobc: make(chan int, 1),
		}
		go job.run()
		n.jobs = append(n.jobs, job)
	}
	n.run()
}
func (n *NumChan) run() {
	for {
		n.seq()
	}
}
func (n *NumChan) seq() {
	for _, j := range n.jobs {
		j.Jobc <- 1
		time.Sleep(time.Second * 1)
	}
}

type Job struct {
	ID   int
	Jobc chan int
}

func (j *Job) run() {
	for {
		select {
		case <-j.Jobc:
			fmt.Printf("id %d\n", j.ID)
		}
	}
}

func main() {
	n := &NumChan{}
	n.JobNum(4)
}
