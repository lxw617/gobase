package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("original a =", a)
	fmt.Printf("初始化 a 的地址=%p \n", &a)
	for i, v := range a {
		fmt.Println("inner for range loop 1111111111111111111111111111")
		fmt.Printf("for range 修改之前 a 的地址=%p \n", &a)
		fmt.Println("before, r =", r)
		fmt.Println("before, a =", a)
		fmt.Println("before, i =", i, " v= ", v)
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
		fmt.Println("after, r =", r)
		fmt.Println("after, a =", a)
	}
	fmt.Printf("最终结果 a 的地址=%p \n", &a)
	fmt.Println("after for range loop, r =", r)
	fmt.Println("after for range loop, a =", a)
}
