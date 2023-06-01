package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"time"

	"github.com/eapache/go-resiliency/retrier"
)

type MessageResponse struct {
	Code    int
	Message string
	Method  string
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

/*
1- Implement function to get log
2- Implement function to validate a http request from client.
3- Using chain handler function in our server
4- Using back off pattern in our server.
5- Using Marshalling and UnMarshalling for our server
*/
func main() {

	// Define port at the begin of program
	port := 8080

	http.HandleFunc("/helloworld", log(HelloWorldHandlerFunc))
	http.HandleFunc("/hello", log(validate(HelloHandlerFunc)))
	fmt.Printf("Server starting on port %v\n", port)

	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)

	// With default server, we don't have any error return
	/*----------------- This comment for learning -----------------------
	if errServer != nil {
		panic(errServer.Error())
	}
	--------------------------------------------------------------------*/

	n := 0
	r := retrier.New(retrier.ConstantBackoff(3, 1*time.Second), nil)

	err := r.Run(func() error {
		fmt.Println("Attempt: ", n)
		n++
		return fmt.Errorf("Failed")
	})

	if err != nil {
		fmt.Println(err)
	}
}

func HelloWorldHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world\n")
}

// func log() is internal function, returns logging for work flow on program.
func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

func HelloHandlerFunc(w http.ResponseWriter, r *http.Request) {

	// Mistake 1: We can use ParseForm() because our server
	fmt.Fprintln(w, r.ParseForm())
	fmt.Fprintln(w, "Hello ", r.Form["name"])

	// Mistake 2: We can't also use read body.
	// len := r.ContentLength
	// body := make([]byte, len)
	// r.Body.Read(body)
	// fmt.Fprintln(w, string(body))

}

// func valivate() is internal function, returns a message and error to user
func validate(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var request helloWorldRequest
		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&request)

		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		h(w, r)
	}
}
