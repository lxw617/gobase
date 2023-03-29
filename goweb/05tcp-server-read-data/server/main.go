package main

import (
	"fmt"
	"log"
	"net"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
)

func main() {
	// 指定服务器 通信协议、IP地址、port。 创建一个用于监听的 socket
	listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Fatal("Error starting tcp server : ", err)
	}
	defer listener.Close()
	fmt.Println("服务器等待客户端建立连接...")
	log.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// 阻塞监听客户端连接请求, 成功建立连接，返回用于通信的socket
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}
		fmt.Println("服务器与客户端成功建立连接！！！")
		// 读取客户端发送的数据
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		//_, err = conn.Write(buf[:n])                    // 读多少写多少。原封不动
		_, err = conn.Write([]byte("Yes, i am ready!\n")) // 读多少写多少。原封不动
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		// 处理数据打印
		fmt.Println("服务器读到数据：", string(buf[:n]))
	}
}
