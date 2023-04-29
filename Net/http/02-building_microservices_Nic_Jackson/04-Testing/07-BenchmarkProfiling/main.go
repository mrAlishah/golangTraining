package main

import (
	"04-Testing/handlers"
	"log"
	"net/http"

	_ "net/http/pprof"
)

func main() {
	//
	//http://localhost:6060/debug/pprof/
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	handler := handlers.Search{}
	err := http.ListenAndServe(":8323", &handler)
	if err != nil {
		log.Fatal(err)
	}
}
