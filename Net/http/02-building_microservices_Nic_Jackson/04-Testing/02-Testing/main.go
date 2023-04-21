package main

import (
	"golangTraining/Net/http/02-building_microservices_Nic_Jackson/04-Testing/01-Testing/handlers"
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
