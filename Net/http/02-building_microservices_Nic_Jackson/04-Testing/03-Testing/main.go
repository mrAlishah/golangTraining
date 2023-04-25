package main

import (
	"03-Testing/handlers"
	"log"
	"net/http"
)

func main() {
	handler := handlers.Search{}
	err := http.ListenAndServe(":8323", &handler)
	if err != nil {
		log.Fatal(err)
	}
}
