package main

import "fmt"

func main() {
	mapa := make(map[string][]int, 0)
	mapa["1"] = []int{1}
	mapa["2"] = []int{1, 2}
	mapa["3"] = []int{1, 2, 3}
	mapa["4"] = []int{1, 2, 3, 4}
	for _, a := range mapa {
		fmt.Println(a)
	}
}
