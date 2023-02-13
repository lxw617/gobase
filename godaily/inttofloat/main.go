package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	totalAmount := 1
	number, _ := new(big.Float).SetPrec(uint(1024)).SetString(strconv.Itoa(totalAmount))

	denominator := big.NewFloat(100)
	denominator1 := number.Quo(number, denominator)
	fmt.Println(denominator1)
}
