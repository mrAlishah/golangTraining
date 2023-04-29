package main

import (
	"06-Testing/data"
	"06-Testing/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var store *data.MongoStore

func main() {
	serverURI := "localhost"
	if os.Getenv("DOCKER_IP") != "" {
		serverURI = os.Getenv("DOCKER_IP")
	}

	store = waitForDB()
	clearDB()
	setupData()

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

func waitForDB() *data.MongoStore {
	for i := 0; i < 10; i++ {
		store, err := data.NewMongoStore("mongodb")
		if err == nil {
			return store
		}

		fmt.Println("Waiting for DB Connection")
		time.Sleep(1 * time.Second)
	}

	return nil
}

func clearDB() {
	store.DeleteAllKittens()
}

func setupData() {
	store.InsertKittens(
		[]data.Kitten{
			data.Kitten{
				Id:     "1",
				Name:   "Felix",
				Weight: 12.3,
			},
			data.Kitten{
				Id:     "2",
				Name:   "Fat Freddy's Cat",
				Weight: 20.0,
			},
			data.Kitten{
				Id:     "3",
				Name:   "Garfield",
				Weight: 35.0,
			},
		})
}
