package main

// static web server

import (
	"fmt"
	"log"
	"net/http"
)

// hello handle function
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// checking the correctness of path
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	//checking the method of request
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
	}

	fmt.Fprintf(w, "Hello!")
}

// form handle function
func formHandler(w http.ResponseWriter, r *http.Request) {
	// checking for error when parse form
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v\n", err)
	}
	fmt.Fprintf(w, "POST request successful\n")

	//getting values from form
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s \n", name)
	fmt.Fprintf(w, "Address = %s \n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server on port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
