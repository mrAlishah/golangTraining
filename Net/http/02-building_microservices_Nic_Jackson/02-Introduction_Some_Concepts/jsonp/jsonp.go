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

/*
Request:
	GET /helloworld?callback=hello
Response:
	hello({"message":"Hello World"})
*/

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "HelloWorld"}
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops")
	}

	//if there is a callback parameter in the query string. This would be provided by the client
	//and indicates the function they expect to be called when the response is returned
	callback := r.URL.Query().Get("callback")
	if callback != "" {
		r.Header.Set("Content-Type", "application/javascript")
		//To return our response in JSONP format all we need to do is wrap the standard response to a JavaScript function call.
		//we are taking the callback function name that was passed by the client and encapsulating the response we would normally send.
		fmt.Fprintf(w, "%s(%s)", callback, string(data))
	} else {
		fmt.Fprint(w, string(data))
	}
}
