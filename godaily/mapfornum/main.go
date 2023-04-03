package main

import "fmt"

func main() {
	// 根据指定数组生成不同三位数字组成的map集合
	var a []int = []int{1, 2, 3, 4, 4, 3, 3, 2, 1}
	mapa := make(map[int]int, 0)
	for _, x := range a {
		for _, y := range a {
			for _, z := range a {
				mapa[x*100+y*10+z] = x*100 + y*10 + z
			}
		}
	}

	// for k, v := range mapa {
	// 	fmt.Printf("%d ---> %d\n", k, v)
	// }

	// if item, ok := mapa[344]; ok {
	// 	fmt.Printf("item:%d", item)
	// }

	item := mapa[344]
	if item != 0 {
		fmt.Printf("item:%d", item)
	}
}
