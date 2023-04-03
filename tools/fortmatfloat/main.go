package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(FormatFloat(float64(200)/100, -1))
	fmt.Println(FormatFloat(float64(2.12145453)/100, -1))
	fmt.Println(FormatFloat(float64(2.12145453)/100, 3))
}

func FormatFloat(num float64, prec int) float64 {
	s := strconv.FormatFloat(num, 'f', prec, 64)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return float64(-1)
	}
	return f
}
