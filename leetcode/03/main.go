package main

import (
	"fmt"
)

func main() {
	s := "pwwkew"
	i := lengthOfLongestSubstring(s)
	fmt.Println(i)
}
func lengthOfLongestSubstring(s string) int {
	l, r, max := 0, 0, 0
	m := map[byte]int{}
	for ; r < len(s); r++ {
		v := s[r]
		if _, ok := m[v]; !ok {
			m[v] = r
		} else {
			//更新r，设置l=之前的r+1
			if m[v]+1 >= l {
				l = m[v] + 1
			}
			m[v] = r
		}
		if r-l+1 > max {
			max = r - l + 1
		}
		fmt.Println(m)
	}
	return max
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
