package main

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
)

const (
	CONN_HOST      = "localhost"
	CONN_PORT      = "8080"
	ADMIN_USER     = "admin"
	ADMIN_PASSWORD = "admin"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func BasicAuth(handler http.HandlerFunc, realm string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 使用r.BasicAuth()获取请求授权头中提供的用户名和密码，然后将其与程序中声明的常量进行比较。如果凭证匹配，则返回处理程序，否则设置WWW-Authenticate和状态码401并在 HTTP 响应流上写入You are Unauthorized to access the application。
		// 最后，我们在main()方法中引入了一个从HandleFunc调用BasicAuth的变更，如下所示：
		user, pass, ok := r.BasicAuth()
		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(ADMIN_USER)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(ADMIN_PASSWORD)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			_, err := w.Write([]byte("You are Unauthorized to access the application.\n"))
			if err != nil {
				log.Fatal("error helloWorld : ", err)
				return
			}
			return
		}
		handler(w, r)
	}
}

func main() {
	// 解析 c, err := base64.StdEncoding.DecodeString(auth[len(prefix):]) 中的 auth[len(prefix):]
	fmt.Println("Basic YWRtaW46YWRtaW4="[len("Basic"):])

	// 传递一个BasicAuth处理程序，而不是nil或DefaultServeMux来处理 URL 模式为/的所有传入请求。
	http.HandleFunc("/", BasicAuth(helloWorld, "Please enter your username and password"))
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
