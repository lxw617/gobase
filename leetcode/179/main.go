package main

import "fmt"

func main() {
	nums := []int{10, 2, 9, 4, 6}
	l := largestNumber(nums)
	fmt.Println(l)
}
func largestNumber(nums []int) string {
	str := ""
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	fmt.Println(nums)
	return str
}
