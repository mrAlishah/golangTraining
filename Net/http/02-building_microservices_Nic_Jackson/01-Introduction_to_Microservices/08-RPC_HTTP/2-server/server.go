package server

import (
	contract "01-Introduction_to_Microservices/08-RPC_HTTP/1-contract"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

const port = 1234

func StartServer() {
	// As same as net/http package, we need to create a handler for RPC-base API
	helloWorld := &HelloWorldHandler{}

	// we have a struct field with methods on it we can register this with the RPC server, don't conform about interface
	/***** net/http package: *****/
	// http.Handle("/helloworld", &helloWorld)
	/***** RPC-base API: *****/
	rpc.Register(helloWorld)

	/*#### NEW 1 ::: FOR RPC as a HTTP Transport protocol ####*/
	//we are calling the rpc.HandleHTTP method, this is a requirement using HTTP with RPC
	//as it will register the HTTP handlers we mentioned earlier with the DefaultServer method.
	rpc.HandleHTTP()

	// In net/http package, we use DefaultMuxServe using http protocol
	/***** net/http package: *****/
	// http.ListenAndServe(fmt.Sprintf(":%v", port), nil)

	// In RPC-base API, we can select the optional protocol for our application.
	/***** RPC-base API: *****/
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}

	log.Printf("Server starting on port %v\n", port)

	/*#### NEW 2 ::: FOR RPC as a HTTP Transport protocol ####*/
	log.Fatal(http.Serve(l, nil))
	//http.Serve(l, nil)
}

/*************************************************************************************************/
/* Handler: Reply message response
/*************************************************************************************************/
// Wrong   sentence: Declare struct type for handler is the same using net/http package.
// Correct sentence: The struct type declaration for the handler is the same as when using net/http package

type HelloWorldHandler struct{}

// In net/http packge, we must use the default: http.ResponseWriter, http.Request for handler.
/***** net/http package: *****/
// func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello")
// }

// In RPC-base API, we can declare any request and response message at contract entity.
/***** RPC-base API: *****/
//func (h *HelloWorldHandler) HelloWorld( reply *contract.HelloWorldResponse, args *contract.HelloWorldRequest) error {

func (h *HelloWorldHandler) HelloWorld(req *contract.HelloWorldRequest, res *contract.HelloWorldResponse) error {
	res.Message = "Hello " + req.Name
	return nil
}
