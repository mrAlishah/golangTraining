package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "Hello " + request.Name}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func main() {
	port := 8080

	// Use FileServer to access to our asset.
	// Note: Our Filesystem is started this application.
	cathandler := http.FileServer(http.Dir("./images"))
	// Use StripPrefix to remove the given prefix from the request URL's path and then invoking h handler
	//In the preceding example, we are registering a StripPrefix handler with our path /cat/. If we did not do this, then the
	//FileServer handler would be looking for our image in the images/cat directory.
	http.Handle("/cat/", http.StripPrefix("/cat/", cathandler))

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
