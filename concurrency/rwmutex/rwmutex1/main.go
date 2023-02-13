package main

import "fmt"

func foo() {
	fmt.Println("in foo")
}

func goo() {
	var i int
	fmt.Println("in goo", i)
}

func main() {
	go goo()
	go foo()
	select {}
}
