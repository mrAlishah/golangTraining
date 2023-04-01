package main

import (
	"log"
	"net/http"
	"time"
)

func time2Handler(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
	return http.HandlerFunc(fn)
}

/*
or
func time2Handler(format string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	})
}
or
func timeHandler(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
}
*/

// https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go
func main() {
	mux := http.NewServeMux()

	th := time2Handler(time.RFC1123)
	mux.Handle("/time", th)

	log.Print("Listening...")
	http.ListenAndServe(":3000", mux)
}

/*
The timeHandler() function now has a subtly different role. Instead of coercing the function
into a handler (like we did previously), we are now using it to return a handler.
There's two key elements to making this work.

First it creates fn, an anonymous function which accesses — or closes over — the format variable
forming a closure. Regardless of what we do with the closure it will always be able to access
the variables that are local to the scope it was created in — which in this case means it'll always
have access to the format variable.

Secondly our closure has the signature func(http.ResponseWriter, *http.Request). As you may remember
from a moment ago, this means that we can convert it into a http.HandlerFunc type
(so that it satisfies the http.Handler interface). Our timeHandler() function then returns
this converted closure.

In this example we've just been passing a simple string to a handler. But in a real-world application
you could use this method to pass database connection, template map, or any other application-level
context. It's a good alternative to using global variables, and has the added benefit of making neat
self-contained handlers for testing.
*/
