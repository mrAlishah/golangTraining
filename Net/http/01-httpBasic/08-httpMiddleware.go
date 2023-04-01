/*
We want to create some middleware which
a) checks for the existence of a Content-Type header and
b) if the header exists, check that it has the mime type application/json.
If either of those checks fail, we want our middleware to write an error message
and to stop the request from reaching our application handlers.
*/
package main

import (
	"fmt"
	"log"
	"mime"
	"net/http"
)

func enforceJSONHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hi there,")

		contentType := r.Header.Get("Content-Type")
		fmt.Fprintln(w, "contentType is"+contentType)

		if contentType != "" {
			mt, _, err := mime.ParseMediaType(contentType)
			if err != nil {
				http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
				return
			}

			if mt != "application/json" {
				http.Error(w, "Content-Type header must be application/json", http.StatusUnsupportedMediaType)
				return
			} else {
				fmt.Fprintln(w, "Header Content-Type: application/json")
			}
		}

		next.ServeHTTP(w, r)
	})
}

func final1(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("OK, Look at console for results"))
	fmt.Fprintln(w, "OK, Look at console for results")
}

func main() {
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final1)
	mux.Handle("/", enforceJSONHandler(finalHandler))

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
