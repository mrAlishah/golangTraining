package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

// https://www.alexedwards.net/blog/making-and-using-middleware
/*
In a trivial case like this our code is fairly clear. But what happens if we want to use this
as part of a larger middleware chain? We could easily end up with a declaration looking something
like this...

http.Handle("/", handlers.LoggingHandler(logFile, authHandler(enforceJSONHandler(finalHandler))))
... And that's pretty confusing!
*/
func main() {
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final3)
	mux.Handle("/", handlers.LoggingHandler(logFile, finalHandler))

	log.Print("Listening on :3000...")
	err = http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

func final3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
