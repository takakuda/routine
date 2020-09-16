package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "GET hello!\n")
	case http.MethodPost:
		fmt.Fprintf(w, "POST hello!\n")
	default:
		fmt.Fprintf(w, "Method not allowed.\n")
	}
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
