package main

import (
	"errors"
	"fmt"
)

// 函数去读取以配置文件init.conf的信息
// 如果文件名传入不正确,我们就返回一个自定文的错误
func readconf(name string) (err error) {
	if name == "config.ini" {
		// 读取..。
		return nil
	} else {
		// 必回一个目足又错误
		return errors.New("读取文件错误.")
	}
}

func test02() {
	err := readconf("config2.ini")
	if err != nil {
		// 如果读取文件发送错误，就输出这个错误,并终止程序
		panic(err)
		fmt.Println("teste2()继续执行....")
	}
}

func main() {
	// 测试自定义错误的使用
	test02()
	fmt.Println("main()下面的代码..")
}
