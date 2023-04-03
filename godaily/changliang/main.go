package main

import "reflect"

func main() {
	// 常量声明的同底类型不同类型变量可以相加

	type myInt int

	const a myInt = 3

	const b int = 2

	const c = int(a) + b

	println(c) // 5

	reflect.TypeOf(c) // int

	// 变量声明的同底类型不同类型变量不可以相加

	// type myInt int

	// var a myInt = 3

	// var b int = 2

	// const c = a + b
}
