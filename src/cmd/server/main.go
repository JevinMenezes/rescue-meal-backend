package main

import (
	"log"
	"net/http"
)

const message = "Go HTTP server running"

/*
@author Jevin
@go-version 1.13.1
*/
func main() {
	mux := http.NewServeMux() //@future replace with https://github.com/gorilla/mux
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", "rmgo")
		w.Header().Set("X-Powered-By", "Jevin")
		w.Write([]byte(message))
	})

	log.Println("server starting")

	err := http.ListenAndServe(":8080", mux) //@upgrade replace with ListenAndServeTLS for https requests

	//error handling if server fails to start
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
