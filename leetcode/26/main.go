package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	l := removeDuplicates(nums)
	fmt.Println(l)
}
func removeDuplicates(nums []int) int {
	hashSet := make(map[int]struct{})
	for _, v := range nums {
		hashSet[v] = struct{}{}
	}
	return len(hashSet)
}
