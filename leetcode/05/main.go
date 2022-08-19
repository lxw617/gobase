package main

import "fmt"

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}
func longestPalindrome(s string) string {
	left, right, max, l, r := 0, 1, 0, 0, 0

	for i := 0; i < len(s); i++ {
		if s[left] == s[right] {
			if left-1 >= 0 {
				left--
				right++
			}
		} else {
			left++
			right++
		}
		if right-left+1 > max {
			max = right - left + 1
			l = left
			r = right + 1
		}
	}
	ss := s[l:r]
	return ss
}
