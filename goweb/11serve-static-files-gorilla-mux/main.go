package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

type Person struct {
	Id   string
	Name string
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	person := Person{Id: "1", Name: "Foo"}
	parsedTemplate, _ := template.ParseFiles("./goweb/11serve-static-files-gorilla-mux/templates/first-template.html")
	err := parsedTemplate.Execute(w, person)
	if err != nil {
		log.Printf("Error occurred while executing the template or writing its output : %v", err)
		return
	}
}

func main() {
	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./goweb/11serve-static-files-gorilla-mux/static/"))))
	router.HandleFunc("/", renderTemplate).Methods("GET")
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
