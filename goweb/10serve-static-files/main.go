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

type Person struct {
	Name string
	Age  string
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	person := Person{Age: "1", Name: "Foo"}
	parsedTemplate, _ := template.ParseFiles("./goweb/10serve-static-files/templates/first-template.html")
	err := parsedTemplate.Execute(w, person)
	if err != nil {
		log.Printf("Error occurred while executing the template or writing its output : %v", err)
		return
	}
}

func main() {
	// 静态文件处理
	fileServer := http.FileServer(http.Dir("./goweb/10serve-static-files/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	// html文件处理
	// http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("../views/pages"))))
	http.HandleFunc("/", renderTemplate)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
