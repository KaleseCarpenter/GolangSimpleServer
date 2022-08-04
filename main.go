package main

import (
	"fmt"
	"log"
	"net/http"
)

// CREATE FORM HANDLER

func formHandler(w http.ResponseWriter, r *http.Request) {
	// when users submit something in the html form as a POST REQUEST
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	// Otherwise send this response to user
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// CREATE HELLO HANDLER -- response and request to server

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	// Just want to print "Hello" - don't want any them to get anything because typing in the browser is a Get method
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	// Otherwise send this response to user
	fmt.Fprintf(w, "Hello! Welcome to my first Golang Site")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer) // this will serve the index.html file
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler) // this will print out Hello to the screen

	fmt.Printf("Starting server at port 8080\n")

	// THE ListenAndServe PACKAGE WILL CREATE THE SERVER
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)

	}
}
