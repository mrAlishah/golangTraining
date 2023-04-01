package main

import (
	"log"
	"net/http"
	"time"
)

func time1Handler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

// https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go
func main() {
	mux := http.NewServeMux()

	// Convert the timeHandler function to a http.HandlerFunc type.
	th := http.HandlerFunc(time1Handler)

	// And add it to the ServeMux.
	mux.Handle("/time", th)

	log.Print("Listening...")
	http.ListenAndServe(":3000", mux)
}
