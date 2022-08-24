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
	// 获取指定字符串汉字个数
	s := "hello库珀"
	count := HanCounter(s)
	fmt.Println(count)
}
