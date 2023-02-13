package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var n int64 = 10
	fmt.Printf("n 的类型 %T，占用字节数为 %d", n, unsafe.Sizeof(n))
	fmt.Println()

	//十进制数形式:如:5.12  .512(必须有小数点)
	num6 := 5.12
	num7 := .1231 //=> 0.123
	fmt.Println("num6=", num6, "num7=", num7)
	// 科学计数法形式
	num8 := 5.1234e2   //?5.1234 *10的2次方
	num9 := 5.1234e2   //?5.1234*10的2次方shift+alt+向下的箭头
	num10 := 5.1234e-2 //?5.1234/10的2次方0.051234
	fmt.Println("num8=", num8, "num9=", num9, "num10=", num10)

	var c1 byte = 'a'
	var c2 byte = 'e' //字符的
	//当我们直接输出byte值,就是输出了的对应的字符的码值// 'a'==>
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	//如果我们希望输出对应字符,需要使用格式化输出fmt.Printf("c1=%c c2=%c\n", c1，c2)
	//var c3 byte = '北'//overflow溢出
	var c3 int = '北' //overflow溢出
	fmt.Printf("c3=%c c3对应码值=%d", c3, c3)
	fmt.Println()

	var c4 int = 22269 //可以直接给某个变量赋一个数字，然后按格式化输出时%c，会输出该数字对应的unicode字符
	fmt.Printf("c3=%c c3对应码值=%d", c4, c4)
	fmt.Println()

	// 字符类型是可以进行运算的
	var n1 = 10 + 'a'
	fmt.Println("n1=", n1)
}
