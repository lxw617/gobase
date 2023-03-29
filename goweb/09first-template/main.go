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

// Person 定义了一个具有Id和Name字段的person结构类型
type Person struct {
	Id   string
	Name string
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	// 初始化一个person结构类型
	person := Person{Id: "1", Name: "Foo"}
	parsedTemplate, _ := template.ParseFiles("./goweb/09first-template/templates/first-template.html")
	err := parsedTemplate.Execute(w, person)
	if err != nil {
		log.Printf("Error occurred while executing the template or writing its output : %v", err)
		return
	}
}

func main() {
	http.HandleFunc("/", renderTemplate)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
