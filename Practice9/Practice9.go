package main

import (
	"html/template"
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func fruitHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("fruit.html")
	check(err)

	err = html.Execute(writer, nil)
	check(err)
}

func meatHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("meat.html")
	check(err)

	err = html.Execute(writer, nil)
	check(err)
}

func main() {
	http.HandleFunc("/fruit", fruitHandler)
	http.HandleFunc("/meat", meatHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
