package main

import (
	"log"
	"net/http"

	"github.com/goji/httpauth"
)

func main() {
	authHandler := httpauth.SimpleBasicAuth("test", "test")

	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final2)
	mux.Handle("/", authHandler(finalHandler))

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

func final2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
