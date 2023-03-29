package main

import (
	"html/template"
	"log"
	"net/http"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func login(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("./goweb/12html-form/templates/login-form.html")
	_ = parsedTemplate.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", login)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
