package main

import (
	"06-Testing/data"
	"06-Testing/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	serverURI := "localhost"
	if os.Getenv("DOCKER_IP") != "" {
		serverURI = os.Getenv("DOCKER_IP")
	}

	store, err := data.NewMongoStore(serverURI)
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.Search{DataStore: store}
	err = http.ListenAndServe(":8323", &handler)
	if err != nil {
		log.Fatal(err)
	}

}
