package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Hello World!")
	if err != nil {
		log.Fatal("error helloWorld : ", err)
		return
	}
}

func main() {
	// 我们将创建一个带有单个处理程序的 HTTP 服务器，它将编写 Hello World！在 HTTP 响应流上，使用 GorillaCompressHandler，以.gzip格式将所有响应发送回客户端。
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT,
		handlers.CompressHandler(mux))
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
