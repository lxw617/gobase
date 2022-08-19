package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	target := 6
	index := twoSum(nums, target)
	fmt.Println(index)
}
func twoSum(nums []int, target int) []int {
	var index []int
	var res []int
	hashSet := make(map[int]struct{})
	for k1 := 0; k1 < len(nums); k1++ {
		for k2 := 0; k2 < len(nums); k2++ {
			num1 := nums[k1]
			num2 := nums[k2]
			if num1+num2 == target && k1 != k2 {
				index = append(index, k1)
				index = append(index, k2)
			}
			for _, v := range index {
				hashSet[v] = struct{}{}
			}
		}
	}
	for k, _ := range hashSet {
		res = append(res, k)
	}
	return res
}
