package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

var GetRequestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World!"))
		if err != nil {
			fmt.Println(err)
		}
	},
)

var PostRequestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("It's a Post Request!"))
		if err != nil {
			fmt.Println(err)
		}
	},
)

var PathVariableHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		_, err := w.Write([]byte("Hi " + name))
		if err != nil {
			fmt.Println(err)
		}
	},
)

func main() {
	router := mux.NewRouter()
	// router.Handle("/", GetRequestHandler).Methods("GET")
	// router.Handle("/post", PostRequestHandler).Methods("POST")
	// router.Handle("/hello/{name}", PathVariableHandler).Methods("GET", "PUT")

	// 加日志
	logFile, err := os.OpenFile("./goweb/08http-server-gorilla-mux-routing/server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
	// 用一个 Gorilla 日志处理程序包装GetRequestHandler，并将一个标准输出流作为写入程序传递给它，这意味着我们只是要求在控制台上以 Apache 通用日志格式用 URL 路径/记录每个请求。
	// router.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(GetRequestHandler))).Methods("GET")
	router.Handle("/", handlers.LoggingHandler(logFile, GetRequestHandler)).Methods("GET")
	router.Handle("/post", handlers.LoggingHandler(logFile, PostRequestHandler)).Methods("POST")
	router.Handle("/hello/{name}", handlers.CombinedLoggingHandler(logFile, PathVariableHandler)).Methods("GET")

	err = http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
