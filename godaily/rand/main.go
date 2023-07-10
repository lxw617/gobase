package main

import (
	cryptorand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	s := fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(100000000))
	fmt.Println(s)

	// 我们一般使用系统时间的不确定性来进行初始化
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		fmt.Print(rand.Intn(10), " ")
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(128)
	fmt.Println(num)

	// 对安全性要求比较高，需要使用真随机数的话，那么可以使用 crypto/rand 包中的方法,这样生成的每次都是不同的随机数.
	// 生成 10 个 [0, 128) 范围的真随机数。
	for i := 0; i < 10; i++ {
		result, _ := cryptorand.Int(cryptorand.Reader, big.NewInt(128))
		fmt.Println(result)
	}
}
