package main

import (
	"bytes"
	"fmt"
)

func main() {
	sqlBuffer := new(bytes.Buffer)
	sqlBuffer.WriteString(`'{"batchNo":"221215213849059326","description":"订单状态错误","isSuccess":0}'`)
	fmt.Println(sqlBuffer.String())
}
