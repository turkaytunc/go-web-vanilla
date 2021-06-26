package main

import (
	"log"
	"net/http"
)

func setupHandler(w http.ResponseWriter, r *http.Request) {

	message := []byte("Hello World")
	_, err := w.Write(message)

	if err != nil {
		log.Fatalf("%s", err)
	}
}

func main() {
	http.HandleFunc("/", setupHandler)
	http.ListenAndServe("localhost:4000", nil)
}
