package main

import (
	"fmt"
	"log"
	"net/http"
)

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("---------------------------")
		log.Print("Executing middlewareOne")
		fmt.Fprintln(w, "Executing middlewareOne")

		next.ServeHTTP(w, r)

		log.Print("Executing middlewareOne again")
		fmt.Fprintln(w, "Executing middlewareOne again")
	})
}

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing middlewareTwo")
		fmt.Fprintln(w, "Executing middlewareTwo")

		if r.URL.Path == "/foo" {
			return
		}
		next.ServeHTTP(w, r)

		log.Print("Executing middlewareTwo again")
		fmt.Fprintln(w, "Executing middlewareTwo again")
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Print("Executing finalHandler")
	fmt.Fprintln(w, "Executing finalHandler")

	w.Write([]byte("OK, See results at the console\n"))
}

// https://www.alexedwards.net/blog/making-and-using-middleware
// to Run as a Client:
// curl -i localhost:3000
// curl -i -H "Content-Type: application/xml" localhost:3000
// curl -i -H "Content-Type: application/json; charset=UTF-8" localhost:3000
func main() {
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", middlewareOne(middlewareTwo(finalHandler)))

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

/*
You'll notice that this middleware function has a func(http.Handler) http.Handler signature.
It accepts a handler as a parameter and returns a handler. This is useful for two reasons:

Because it returns a handler we can register the middleware function directly
with the standard http.ServeMux router in Go's net/http package.
We can create an arbitrarily long handler chain by nesting middleware functions inside each other.
For example:

mux := http.NewServeMux()
mux.Handle("/", middlewareOne(middlewareTwo(finalHandler)))
*/
