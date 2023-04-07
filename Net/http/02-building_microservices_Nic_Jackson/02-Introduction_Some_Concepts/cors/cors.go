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

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
	log.Fatal("Something")

}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		/* we detect if the request method is OPTIONS and instead of returning the response
		we return the Access-Control-Allow-Origin header that the client is expecting.
		"*" : which means all domains are allowed to interact with this API. This is not the safest implementation
		and quite often you will request your API users to register the domains that will be interacting
		with the API and restrict the Allow-Origin to only include those domains.
		we have hard coded this into the handler.
		*/
		w.Header().Add("Access-Control-Allow-Origin", "*")
		//This tells the browser that it can only make GET requests to this URI and that it is forbidden to make POST, PUT, and so on.
		w.Header().Add("Access-Control-Allow-Methods", "GET")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	response := helloWorldResponse{Message: "Hello World"}
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops")
	}

	fmt.Fprint(w, string(data))
}
