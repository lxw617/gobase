package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

var cookieHandler *securecookie.SecureCookie

func init() {
	cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))
}

func createCookie(w http.ResponseWriter, r *http.Request) {
	value := map[string]string{
		"username": "Foo",
	}
	base64Encoded, err := cookieHandler.Encode("key", value)
	if err == nil {
		cookie := &http.Cookie{
			Name:  "first-cookie",
			Value: base64Encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
	_, err = w.Write([]byte("Cookie created."))
	if err != nil {
		log.Fatal("error : ", err)
	}
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	log.Printf("Reading Cookie..")
	cookie, err := r.Cookie("first-cookie")
	if cookie != nil && err == nil {
		value := make(map[string]string)
		if err = cookieHandler.Decode("key", cookie.Value, &value); err == nil {
			_, err = w.Write([]byte(fmt.Sprintf("Hello %v \n", value["username"])))
			if err != nil {
				log.Fatal("error : ", err)
			}
		}
	} else {
		log.Printf("Cookie not found..")
		_, err = w.Write([]byte("Hello"))
		if err != nil {
			log.Fatal("error : ", err)
		}
	}
}

func main() {
	http.HandleFunc("/create", createCookie)
	http.HandleFunc("/read", readCookie)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
