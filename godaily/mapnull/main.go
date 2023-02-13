package main

func main() {

	var m = make(map[*int]*int)
	m[nil] = nil

	println(m)
	println(len(m))

	var b = 5
	m[nil] = &b
	p := m[nil]

	println(*p)
}

// 0xc0000466b0

// 1

// 5
