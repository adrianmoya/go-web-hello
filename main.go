package main

import (
	"adrianmoya.com/go-web-hello/Godeps/_workspace/src/github.com/gorilla/mux"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Printf("Using port %s", port)

	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love Go (or what?)!")
}
