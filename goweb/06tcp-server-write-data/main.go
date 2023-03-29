package main

import (
	"fmt"
	"net"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
)

func main() {
	// 指定 服务器 IP + port 创建 通信套接字。
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()
	// 主动写数据给服务器
	_, err = conn.Write([]byte("Hello to TCP server\n "))
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	buf := make([]byte, 4096)
	// 接收服务器回发的数据
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	fmt.Println("服务器回发：", string(buf[:n]))
}
