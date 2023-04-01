package main

import (
	"log"
	"net/http"
)

// https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go
func main() {
	// Use the http.NewServeMux() function to create an empty servemux.
	mux := http.NewServeMux()

	// Use the http.RedirectHandler() function to create a handler which 307
	// redirects all requests it receives to http://example.org.
	rh := http.RedirectHandler("http://example.org", 307)

	// Next we use the mux.Handle() function to register this with our new
	// servemux, so it acts as the handler for all incoming requests with the URL
	// path /foo.
	mux.Handle("/foo", rh)

	log.Print("Listening...")

	// Then we create a new server and start listening for incoming requests
	// with the http.ListenAndServe() function, passing in our servemux for it to
	// match requests against as the second parameter.
	http.ListenAndServe(":3000", mux)
}

/*
Processing HTTP requests with Go is primarily about two things: handlers and servemuxes.

If you’re coming from an MVC-background, you can think of handlers as being a bit like controllers.
Generally speaking, they're responsible for carrying out your application logic and writing response
headers and bodies.

Whereas a servemux (also known as a router) stores a mapping between the predefined URL paths
for your application and the corresponding handlers. Usually you have one servemux for your application
containing all your routes.

Go's net/http package ships with the simple but effective http.ServeMux servemux,
plus a few functions to generate common handlers including http.FileServer(), http.NotFoundHandler()
and http.RedirectHandler().
*/
