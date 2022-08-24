package main

import "fmt"

func main() {
	// nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	nums := []int{1, 1, 2}
	l := removeDuplicates(nums)
	fmt.Println(l)
}

// func removeDuplicates(nums []int) int {
// 	intInSlice := func(i int, list []int) bool {
// 		for _, v := range list {
// 			if v == i {
// 				return true
// 			}
// 		}
// 		return false
// 	}
// 	var Uniq []int
// 	for _, v := range nums {
// 		if !intInSlice(v, Uniq) {
// 			Uniq = append(Uniq, v)
// 		}
// 	}
// 	fmt.Println(Uniq)
// 	return len(Uniq)
// }
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 1, 1
	for right < len(nums) {
		if nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}
