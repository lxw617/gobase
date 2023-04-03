package main

import (
	"fmt"
	"sync"
)

//--------------------------------
type Ticket struct {
	ID    int
	Price string
	Type  string
}

//--------------------------------
type Conductor struct {
	Tickets []Ticket
	Count   int
}

func NewConductor() *Conductor {
	/*tickets := make([]Ticket, 0)
	tickets = append(tickets, Ticket{
		ID:    0,
		Price: "1",
		Type:  "2",
	})
	tickets = append(tickets, Ticket{
		ID:    -1,
		Price: "12",
		Type:  "1",
	})
	return &Conductor{tickets, 0}*/
	return &Conductor{make([]Ticket, 0), 0}
}

func (this *Conductor) Sell() Ticket {
	temp := this.Tickets[0]
	this.Tickets = this.Tickets[1:]

	return temp
}

func (this *Conductor) Product() Ticket {
	ticket := Ticket{this.Count + 1, "$18", "成人票"}
	this.Count = this.Count + 1
	this.Tickets = append(this.Tickets, ticket)
	return ticket
}

func main() {
	conductor := NewConductor()

	ticketCount := 100

	temp := make(chan int, 100)

	wg := new(sync.WaitGroup)
	wg.Add(ticketCount)
	// 一个线程产票，一个线程售票
	for i := 1; i <= ticketCount; i++ {
		go func() {
			defer wg.Done()
			<-temp
			fmt.Println("出售：", conductor.Sell())
		}()
	}
	for i := 1; i <= ticketCount; i++ {
		go func() {
			//	defer wg.Done()
			temp <- i
			fmt.Println("生产：", conductor.Product())
		}()
	}
	wg.Wait()
}
