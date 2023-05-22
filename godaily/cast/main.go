package main

import (
	"fmt"
	"math"
)

func main() {
	cast := int(Round(1/100.00, 0)) // 四舍五入
	fmt.Println(cast)
}
func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}
