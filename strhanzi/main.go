package main

import (
	"fmt"
	"unicode"
)

// HanCounter to count the number of chinese character.
func HanCounter(s string) int {
	var count int // 0
	for _, c := range s {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	return count
}

func main() {
	s := "hello库珀"
	count := HanCounter(s)
	fmt.Println(count)
}
