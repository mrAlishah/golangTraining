# Introduction
This tutorial will help you implement a simple RPC_HTTP with Hello World message. RPC as HTTP Transport protocol 

## Server
### Start Server
```go
func StartServer() {
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld)
	rpc.HandleHTTP() //It's New

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}
    
	log.Printf("Server starting on port %v\n", port)
	
	log.Fatal(http.Serve(l, nil)) //It's New
	
	//Old Version	
	//defer l.Close()
	//
	//for {
	//	conn, _ := l.Accept()
	//	go rpc.ServeConn(conn)
	//}
}
```
### Server Handler
To implement a handler for RPC. Handler must have 2 arguement to apdapt rpc package.
* req || args *contract.HelloWorldRequest: For message request
* res || reply *contract.HelloWorldResponse: For message response
```go
type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}
```
## Client
### Create Client
```go
const port = 1234

func CreateClient() *rpc.Client {
	//err := client.Call("HelloWorldHandler.HelloWorld", req, &res)
	//It's New
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("localhost:%v", port))
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
func PerformRequest(c *rpc.Client) contract.HelloWorldResponse {
	args := &contract.HelloWorldRequest{Name: "World"}
	var reply contract.HelloWorldResponse
	err := c.Call("HelloWorldHandler.HelloWorld", args, &reply)

	if err != nil {
		log.Fatal("error:", err)
	}

	return reply
}
```

# Getting Started
To run rpc example
```go
go run ./main.go
```

To run test
```
cd 08-PRC_HTTP
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
