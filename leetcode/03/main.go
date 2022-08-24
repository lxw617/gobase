package main

import (
	"fmt"
)

func main() {
	s := "pwwkew"
	i := lengthOfLongestSubstring(s)
	fmt.Println(i)
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

/*
l, r, Max := 0, 0, 0
	m := map[byte]int{}
	for ; r < len(s); r++ {
		fmt.Println(m[s[r]])
		fmt.Println(string(s[r]))
		if _, ok := m[s[r]]; !ok {
			m[s[r]] = r
		} else {
			if m[s[r]]+1 >= l {
				l = m[s[r]] + 1
			}
			m[s[r]] = r
		}
		if r-l+1 > Max {
			Max = r - l + 1
		}
	}
	return Max
*/
