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
	log.Fatal(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil))
}
