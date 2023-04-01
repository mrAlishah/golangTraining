package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}

func main() {
	port := 8080

	//The HandleFunc method creates a Handler
	//type on the DefaultServeMux handler, mapping the path passed in the first parameter to the function in the second parameter:
	//func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	// The HandleFunc method creates a Handler type on the DefaultServeMux handler,
	// mapping the path passed in the first parameter to the function in the second parameter.
	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", 8080)

	//ListenAndServe takes two parameters, the TCP network address to bind the server to and
	//the handler that will be used to route requests:
	//func ListenAndServe(addr string, handler Handler) error
	//The second parameter we are passing is nil, this is because we are using the DefaultServeMux handler,
	//which we are setting up with our call to http.HandleFunc.
	// Function: ListenAndServe takes two parameters, the TCP network address to bind the server to and the handler that will be used to route requests
	// @Para1: port -- network address 8080 bind the server to all available IP addresses on port 8080.
	// @Para2: nil  -- this is because we are using the DefaultServeMux handler, which we are setting up with our call to http.
	// Since ListenAndServe blocks if the server starts correctly we will never exit on a successful start.
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
	//we are passing the output of ListenAndServe
	//straight to log.Fatal(error), which is a convenience function equivalent to calling fmt.Print(a ...interface{}) followed
	//by a call to os.Exit(1). Since ListenAndServe blocks if the server starts correctly we will never exit on a successful start.
}

/* Error
If you do get this error message it means that you are already running an application on your computer that is using port 8080,
this could be another instance of your program or it could be another application. You can check that you do not have another
instance running by checking the running processes:
$ ps -aux | grep 'go run'
*/
