package main

import (
	"fmt"
	"log"
	"net/http"
)

//you need response and request.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, " GET method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Welcome to the hello world page!!")
}

//you need response and request. For form handler
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseform error : %v", err)
		return
	}

	fmt.Fprint(w, "POST successful")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name  - %s\n", name)
	fmt.Fprintf(w, "Address - %s\n", address)

}

func main() {
	//main logic comes here

	//automatically check static for index.html file
	fileServer := http.FileServer(http.Dir("./static"))

	//ask go to handle base route using fileServer
	http.Handle("/", fileServer)

	//ask go to handle functions
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")

	err := http.ListenAndServe(":8080", nil)

	//handle errors
	if err != nil {
		log.Fatal(err)
	}

}
