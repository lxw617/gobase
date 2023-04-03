package main

func main() {
	m := make(map[*int]*int)
	m[nil] = nil

	println(m)
	println(m[nil])
	println(len(m))

	b := 5
	m[nil] = &b
	p := m[nil]

	println(*p)
	println(p)
}

// 0xc0000466b0
// 1
// 5
