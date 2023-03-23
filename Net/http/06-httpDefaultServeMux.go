package main

import (
	"log"
	"net/http"
	"time"
)

func time3Handler(format string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	})
}

// https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go
func main() {
	// Note that we skip creating the ServeMux because we used global variable
	// var DefaultServeMux = NewServeMux()
	//Additionally, http.ListenAndServe() will fall back to using the default servemux if no other handler is provided (that is, the second parameter is set to nil).

	var format string = time.RFC1123
	th := time3Handler(format)

	// We use http.Handle instead of mux.Handle...
	http.Handle("/time", th)

	log.Print("Listening...")
	// And pass nil as the handler to ListenAndServe.
	http.ListenAndServe(":3000", nil)
}
