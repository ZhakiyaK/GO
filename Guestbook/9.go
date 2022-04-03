package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"test/datafile"
)

type Guestbook struct {
	SignatureCount int
	Signature      []string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html")
	check(err)

	err = html.Execute(writer, nil)
	check(err)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	Signature := request.FormValue("Signature")

	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, Signature)
	check(err)
	err = file.Close()
	check(err)

	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signature, err := datafile.GetStrings("signatures.txt")
	check(err)
	html, err := template.ParseFiles("view.html")
	check(err)

	guestbook := Guestbook{
		SignatureCount: len(signature),
		Signature:      signature,
	}

	err = html.Execute(writer, guestbook)
	check(err)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
