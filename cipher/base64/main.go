package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	src := base64.StdEncoding.EncodeToString([]byte("admin:admin"))
	fmt.Println(string(src))
}
