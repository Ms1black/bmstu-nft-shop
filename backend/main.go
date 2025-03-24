package main

import (
	"net/http"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)
	http.ListenAndServeTLS(":443", "cert.pem", "key.pem", handler)
}

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "BMSTU NFT Store Backend")
// }

// func main() {
// 	http.HandleFunc("/", homeHandler)
// 	log.Fatal(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil))
// }
