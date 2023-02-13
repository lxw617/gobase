package main

import (
	"fmt"

	"github.com/google/go-querystring/query"
)

func main() {
	type Options struct {
		// Query   string `url:"q"`
		ShowAll bool `url:"all"`
		a       bool `url:"a"`
		// Page    int    `url:"page"`
		effectStock bool `url:"effectStock"`
	}

	// opt := Options{"foo", true, 2}
	opt := Options{true, true, true}
	v, _ := query.Values(opt)
	fmt.Print(v.Encode()) // will output: "q=foo&all=true&page=2"
}
