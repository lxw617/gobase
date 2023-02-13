package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	valid := []byte("Hello, 世界")
	invalid := []byte{0xff, 0xfe, 0xfd}
	valid1 := []byte{231, 149, 140} // unicode字符“界”的utf8编码

	fmt.Println(valid)               // [72 101 108 108 111 44 32 228 184 150 231 149 140]
	fmt.Println(invalid)             // [255 254 253]
	fmt.Println(utf8.Valid(valid))   // true
	fmt.Println(utf8.Valid(invalid)) // false
	fmt.Println(valid1)              // [231 149 140]
	fmt.Println(utf8.Valid(valid1))  // true
}
