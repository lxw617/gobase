package main

import "fmt"

func main() {
	// nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	nums := "AAAAAAAAAAAAA"

	l := findRepeatedDnaSequences(nums)
	fmt.Println(l)
}
func findRepeatedDnaSequences(s string) []string {
	arrStr := make([]string, 0)
	for i := 0; i+10 < len(s); i++ {
		str := s[i : i+10]
		for j := i; j+10 < len(s); j++ {
			if i != j {
				innerStr := s[j : j+10]
				if str == innerStr {
					arrStr = append(arrStr, str)
					break
				}
			}
		}
	}
	if len(arrStr) > 1 {
		if arrStr[0] == arrStr[1] {
			arrStr = append(arrStr[:1], arrStr[2:]...)
		}
	}
	return arrStr
}
