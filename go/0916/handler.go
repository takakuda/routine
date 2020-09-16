package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)

		name := r.URL.Query().Get("name")
		fmt.Fprintf(w, "GET hello %s!\n", name)
	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "POST hello!\n")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed.\n")
	}
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
