package main 

import (
	"net/http"
	"fmt"
	"os"
	"log"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"		
	}

	log.Printf("Using port %s", port) 

	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	http.ListenAndServe(":"+port,r)
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hi there, I love Go (or what?)!")
}

