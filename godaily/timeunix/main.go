package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Unix(int64(1682575980), 0))
	fmt.Println(time.Unix(int64(1682575980), 1000))
	fmt.Println(time.Unix(0, int64(1682575980)))
}
