package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 	re := regexp.MustCompile("a(x*)b(y|z)c")
	// 	fmt.Printf("%q\n", regexp.MustCompile("-axxxbyc-")) //["axxxbyc" "xxx" "y"]
	// 	r1 := regexp.MustCompile("p([a-z]+)ch")
	// b1:
	// 	r1.MatchString("peach")
	// 	fmt.Println(b) //结果为true

	r, _ := regexp.Compile("^SD{d}")
	b := r.MatchString("peach")
	fmt.Println(b) //结果为true

}
