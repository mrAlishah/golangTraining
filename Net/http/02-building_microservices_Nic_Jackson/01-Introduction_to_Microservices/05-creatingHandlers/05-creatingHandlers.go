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

// curl localhost:8080/helloworld -d '{"name":"Alishah"}'
// curl -X POST localhost:8080/helloworld -H "Content-Type: application/json" -d '{"name":"Alishah"}'
func main() {
	port := 8080

	/********** Chain Handlers Solutions ********/
	// the first handler: validation handler,
	// the second handler: return response if validation is OK.
	// newValidationHandler(next http.Handler Interface{} as last responsible chain) -> next.ServeHTTP(rw, r)
	// -> invoke helloWorldHandler.ServeHTTP(http.ResponseWriter,*http.Request) -> next chain ...to END
	handler := newValidationHandler(newHelloWorldHandler())

	// Note: we're going to build a handler, not function handler.
	http.Handle("/helloworld", handler)

	/********** Using handler from Go-web-application ********/
	// We need initialize a instance for this struct. We also use constructor to instead of it.
	hello := HelloHandler{}
	// Implement handler by pointer
	http.Handle("/hello", &hello)

	// We need initialize a instance for this struct. We also use constructor to instead of it.
	world := WorldHandler{}
	// Implement handler by value
	http.Handle("/world", world)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

/*************************************************************************************************/
/* Handler 1: Validation Handler. This is ServHTTP first chain
/* First, validationHandler{next http.Handler Interface}.ServeHTTP invoke to read http.Request and validate
/* Then, validationHandler.next.ServeHTTP(rw, r) fire to invoke helloWorldHandler struct{}.ServeHTTP to Response
/*************************************************************************************************/
// Create struct validation for handler
// we need to have a reference as an input Param (next http.Handler Interface{}) to the next in the chain
//as it has the responsibility for calling ServeHTTP or returning a response.
type validationHandler struct {
	next http.Handler
}

// Return constructor object with arguement
// @para: next pass http.Handler to validation.
// @return: http.Handler for validation
func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

// Implement handler for validation.
func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}
	// From h - validation handler, pass ResponseWriter, Request to Validation Handler.
	h.next.ServeHTTP(rw, r)
}

/*************************************************************************************************/
/* Handler 2: Reply message response. This is ServHTTP last chain
/*************************************************************************************************/
// Create struct response for handler
type helloWorldHandler struct{}

// Return constructor for new handler
func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

// Implement handler for response.
func (h helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello"}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

/*************************************************************************************************/
/* Handler: just replay "Hello", It's simple implement Handler if any struct implement ServeHTTP
/* It can be as a http.Handler to use http.Handle("/hello", &hello)
/*************************************************************************************************/

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

/*************************************************************************************************/
/* Handler: just replay "World", It's simple implement Handler from Go-book-application,
/*************************************************************************************************/

type WorldHandler struct{}

func (h WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World")
}
