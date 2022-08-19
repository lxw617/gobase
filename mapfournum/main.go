package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4, 4, 3, 3, 2, 1}
	mapa := make(map[int]int, 0)
	for _, x := range a {
		for _, y := range a {
			for _, z := range a {
				mapa[x*100+y*10+z] = x*100 + y*10 + z
			}
		}
	}
	for k, v := range mapa {
		fmt.Printf("%d ---> %d\n", k, v)
	}
}
