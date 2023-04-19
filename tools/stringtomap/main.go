package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "alipay.Response openapi gateway response error, code: 40004, msg: Business Failed, sub_code: SETTLED_ALIPAYACCOUNT_NOT_EXIST, sub_msg: 支付宝账号非法"
	m := make(map[string]string)
	pairs := strings.Split(str, ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, ":")
		if len(kv) == 2 {
			m[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}
	fmt.Println(m)
	subMsg, ok := m["sub_msg"]
	fmt.Println(subMsg, ok)
}
