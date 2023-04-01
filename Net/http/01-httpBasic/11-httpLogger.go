package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

//https://www.alexedwards.net/blog/making-and-using-middleware
/**
To help tidy this up it's possible to create a constructor function which wraps the LoggingHandler()
middleware and returns a standard func(http.Handler) http.Handler function that
we can nest neatly with other middleware.
Or
The justinas/alice package is very lightweight tool which provides some syntactic sugar
for chaining middleware handlers. At it's most basic, it lets you rewrite this:
https://github.com/justinas/alice
mux.Handle("/", loggingHandler(authHandler(enforceJSONHandler(finalHandler))))
As this:
mux.Handle("/", alice.New(loggingHandler, authHandler, enforceJSONHandler).Then(finalHandler))
*/
func newLoggingHandler(dst io.Writer) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return handlers.LoggingHandler(dst, h)
	}
}

func main() {
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal(err)
	}

	loggingHandler := newLoggingHandler(logFile)

	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final4)
	mux.Handle("/", loggingHandler(finalHandler))

	log.Print("Listening on :3000...")
	err = http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

func final4(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
