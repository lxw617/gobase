package main

import (
	"bufio"
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
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Print("Message Received from the client: ", string(message))
	_, err = conn.Write([]byte(message + "\n"))
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	conn.Close()
}
