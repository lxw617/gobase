package main

func main() {
	type T struct {
		F1 int
		F2 string
		f3 int
		F4 int
		F5 int
	}

	// var t1 := T{11, "hello", 13} // 错误：implicit assignment of unexported field 'f3' in T literal

	// var t3 := T{11, "hello", 13, 14, 15}
	// 错误：implicit assignment of unexported field 'f3' in T literal

}
