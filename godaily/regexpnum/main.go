package main

import (
	"fmt"
	"regexp"
)

func main() {
	strArr := []string{"sascsa", "aaaaaaaaaaa", "21111111111", "1115489", "12345678911", "1@154879645", "111111111111"}
	reg := regexp.MustCompile(`^1\d{10}$`)
	// `^(13[0-9]|14[01456879]|15[0-3,5-9]|16[2567]|17[0-8]|18[0-9]|19[0-3,5-9])\d{8}$`

	// reg := regexp.MustCompile(`[0-9]\d{11}`)
	for _, v := range strArr {
		ok := reg.MatchString(v)
		fmt.Println(ok)
	}

}
