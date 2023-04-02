// Please read to understand about reading_writing_json_7.go, and then we will investigate this part.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
The reason we are not just using a simple string is that context often flows across packages and if we just used string then we
could end up with a key clash where one package within your control is writing a name key and another package which is
outside of your control is also using the context and writing a key called name, in this instance the second package would
inadvertently overwrite your context value. By declaring a package level type validationContextKey and using this we can
ensure that we avoid these collisions.
*/
type validationContextKey string

type helloWorldResponse struct {
	Message string `json:"message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

// curl localhost:8080/helloworld -d '{"name":"Alishah"}'
// curl -X POST localhost:8080/helloworld -H "Content-Type: application/json" -d '{"name":"Alishah"}'
func main() {
	port := 8080

	handler := newValidationHandler(newHelloWorldHandler())
	http.Handle("/helloworld", handler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

/*************************************************************************************************/
/* Handler 1: Validation
/*************************************************************************************************/
type validationHandler struct {
	next http.Handler
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}
	// Step 1: WithValue() method to get the parent Context and associated with key.
	c := context.WithValue(r.Context(), validationContextKey("name"), request.Name)
	// Step 2: The WithContext object returns a shallow copy of the original
	// request which has the context changed to the given ctx context.
	// It contains connection between client and server.
	r = r.WithContext(c)

	h.next.ServeHTTP(rw, r)
}

/*************************************************************************************************/
/* Handler 2: Reply message response
/*************************************************************************************************/
type helloWorldHandler struct {
}

func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

func (h helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	// Step 3: Valiadate request is correctly context by using Value() methods.
	name := r.Context().Value(validationContextKey("name")).(string)

	response := helloWorldResponse{Message: "Hello " + name}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}
