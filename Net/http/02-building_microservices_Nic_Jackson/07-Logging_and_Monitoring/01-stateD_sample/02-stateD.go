package main

import (
	"fmt"
	"gopkg.in/alexcesaro/statsd.v2"
	"log"
	"net/http"
	"os"
	"stateDSample/handlers"
)

const port = 8091

// curl -d '{"name": "world"}' http://localhost:8091/helloworld
// Bad Request: curl  http://localhost:8091/helloworld
func main() {
	statsd, err := createStatsDClient(os.Getenv("STATSD"))
	if err != nil {
		log.Fatal("Unable to create statsD client")
	}

	http.Handle("/helloworld", handlers.NewHelloWorldHandler(statsd))
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func createStatsDClient(address string) (*statsd.Client, error) {
	// it needs Storage and querying as a service for stateD metrics data
	//There are multiple options for storing and querying metric data; you have the possibility for self-hosting, or you can utilize a
	//software as a service.we will implement it in docker-compose by image: prom/statsd-exporter
	//return statsd.New(statsd.Address(address))
	return statsd.New()
}
