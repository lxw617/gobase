package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{999999998, 999999997, 999999999}
	l := largestNumber(nums)
	fmt.Println(l)
}
func largestNumber(nums []int) string {
	str := ""
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			left := strconv.Itoa(nums[j])
			right := strconv.Itoa(nums[j+1])
			if left+right < right+left {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	for i := 0; i <= len(nums)-1; i++ {
		str += strconv.Itoa(nums[i])
	}
	if str[0] == '0' {
		return "0"
	}
	return str
}
