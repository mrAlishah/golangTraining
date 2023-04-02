package server

import (
	contract "01-Introduction_to_Microservices/07-RPC/1-contract"
	"fmt"
	"log"
	"net"
	"net/rpc"
)

const port = 1234

func main() {
	log.Printf("Server starting on port %v\n", port)
	StartServer()
}

func StartServer() {
	// As same as net/http package, we need to create a handler for RPC-base API
	helloWorld := &HelloWorldHandler{}

	// we have a struct field with methods on it we can register this with the RPC server, don't conform about interface
	/***** net/http package: *****/
	// http.Handle("/helloworld", &helloWorld)
	/***** RPC-base API: *****/
	rpc.Register(helloWorld)

	// In net/http package, we use DefaultMuxServe using http protocol
	/***** net/http package: *****/
	// http.ListenAndServe(fmt.Sprintf(":%v", port), nil)

	// In RPC-base API, we can select the optional protocol for our application.
	/***** RPC-base API: *****/
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}

	// Defer, after end process, it will close connection.
	defer l.Close()

	for {
		// To receive connections,  we must call the Accept method on the listener by using Accetp().
		//you will see that we have an endless for loop, this is because unlike ListenAndServe which blocks
		//for all connections, with an RPC server we handle each connection individually and as soon as we
		//deal with the first connection we need to continue to again call Accept to handle subsequent
		//connections or the application would exit.

		/*
			with an RPC server we handle each connection individually and as soon as we
			deal with the first connection we need to continue to again call Accept to handle subsequent connections or the application would exit. Accept is a
			blocking method so if there are no clients currently attempting to connect to the service then Accept will block until one does. Once we receive a
			connection then we need to call the Accept method again to process the next connection.
		*/
		conn, _ := l.Accept()

		// The ServeConn method runs the DefaultServer method on the given connection, and will block until the client completes.
		// because of this we use go
		go rpc.ServeConn(conn)
	}
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

func (h *HelloWorldHandler) HelloWorld(res *contract.HelloWorldResponse, req *contract.HelloWorldRequest) error {
	res.Message = "Hello " + req.Name
	return nil
}
