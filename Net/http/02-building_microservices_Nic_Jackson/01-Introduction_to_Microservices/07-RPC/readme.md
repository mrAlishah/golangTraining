# Introduction
This tutorial will help you implement a simple RPC with Hello World message.  

## Server
### Start Server
```go
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
	defer l.Close()

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}
}

```
### Server Handler
To implement a handler for RPC. Handler must have 2 arguement to apdapt rpc package.
* req || args *contract.HelloWorldRequest: For message request
* res || reply *contract.HelloWorldResponse: For message response
```go 
type HelloWorldHandler struct{}

// In net/http packge, we must use the default: http.ResponseWriter, http.Request for handler.
/***** net/http package: *****/
// func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello")
// }

// In RPC-base API, we can declare any request and response message at contract entity.
/***** RPC-base API: *****/
//func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse,) error {
// req || args
// res || reply

func (h *HelloWorldHandler) HelloWorld(req *contract.HelloWorldRequest, res *contract.HelloWorldResponse) error {
res.Message = "Hello " + req.Name
return nil
}
```
## Client
### Create Client
```go
const port = 1234

func CreateClient() *rpc.Client {
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal("dialing:", err)
	}

	return client
}

```

### Create Request Client
If a client want to perform a request to server, they must follow rules to reply messsage
* `client *rpc.Client`: input parameter is a rpc.Client
* `Call()`: must specific handler message "HelloWorldHandler.HelloWorld"
*  req || args
*  res || reply
```go
func PerformRequest(client *rpc.Client) contract.HelloWorldResponse {
	req := &contract.HelloWorldRequest{Name: "world"}
	var res contract.HelloWorldResponse

	// After connect to server successfully, when we use client.Call(), it can automatically run Dial() or something else
	// Then we can make a request to the Server by inputting parameters into client.Call().
	err := client.Call("HelloWorldHandler.HelloWorld", req, &res)
	if err != nil {
		log.Fatal("error:", err)
	}

	return res
}
```

# Getting Started
To run rpc example
```go
go run ./main.go
```

To run test
```
cd 07-PRC
go test -v -run="none" -bench=. -benchtime="5s"
```

Output
```
Hello World
Hello World
Hello World
Hello World
Hello World
Hello World
Hello World
Hello World
Hello World
Hello World
```
