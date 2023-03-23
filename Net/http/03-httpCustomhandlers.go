/*
The handlers that ship with net/http are useful,
but most of the time when building a web application you'll want to use your own custom handlers instead.
*/
package main

import (
	"log"
	"net/http"
	"time"
)

type timeHandler struct {
	format string
}

// implemented interface ServeHTTP(http.ResponseWriter, *http.Request)
func (th timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}

// https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go
func main() {
	mux := http.NewServeMux()

	// Initialise the timeHandler in exactly the same way we would any normal
	// struct.
	th := timeHandler{format: time.RFC1123}

	// Like the previous example, we use the mux.Handle() function to register
	// this with our ServeMux.
	// because of th implemented interface ServeHTTP(http.ResponseWriter, *http.Request)
	// we can use it as a Handler
	mux.Handle("/time", th)

	log.Print("Listening...")
	http.ListenAndServe(":3000", mux)
}

/*
1- When our Go server receives an incoming HTTP request it hands it off to our servemux
(the one that we passed to the http.ListenAndServe() function).
2- The servemux then looks up the appropriate handler based on the request path
(in this case, the /time path maps to our timeHandler handler).
3- The serve mux then calls the ServeHTTP() method of the handler, which in turn writes
out the HTTP response.
*/
