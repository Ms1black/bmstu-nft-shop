package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "BMSTU NFT Store Backend")
}

func main() {
	http.HandleFunc("/", homeHandler)
	log.Println("Starting HTTPS server on https://localhost:8080")
	err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal("Error starting HTTPS server:", err)
	}
}
