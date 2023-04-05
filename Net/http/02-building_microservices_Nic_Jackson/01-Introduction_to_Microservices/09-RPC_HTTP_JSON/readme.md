# Introduction
we will look at the package that provides a built-in codec for serializing and deserializing to the JSON-RPC standard. We will also look at how we can send these responses over HTTP.
## Server
it is the standard rpc server setup,instead of starting the RPC server. we are starting http server and passing the listener to it along with a handler.

### Start Server
```go
const port = 1234

func StartServer() {
	helloWorld := new(HelloWorldHandler)
	rpc.Register(helloWorld)

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}

	log.Printf("Server starting on port %v\n", port)

	//server we are starting http server instead of rpc server and use http.HandlerFunc
	http.Serve(l, http.HandlerFunc(httpHandler))
}
```
**built-in codec for serializing and deserializing to the JSON-RPC standard.**
we are calling the `jsonrpc.NewServerCodec` function and passing to it a type that implements `io.ReadWriteCloser`.The NewServerCodec method returns a type that implements `rpc.ClientCodec`, which has the `ClientCodec interface` methods.
The NewServerCodec method requires that we pass it a type that implements the `ReadWriteCloser interface`. As we do not have such a type passed to us as parameters in the httpHandler method we have defined our own type,`HttpConn`.
```go
func httpHandler(w http.ResponseWriter, r *http.Request) {
	serverCodec := jsonrpc.NewServerCodec(&HttpConn{in: r.Body, out: w})

	err := rpc.ServeRequest(serverCodec)
	if err != nil {
		log.Printf("Error while serving JSON request: %v", err)
		http.Error(w, "Error while serving JSON request, details have been logged.", 500)
		return
	}
}
```
**io.ReadWriteCloser was implemented for rpc.ClientCodec**
A ClientCodec type implements the writing of RPC request and reading RPC responses. To write a request to the connection a client calls the WriteRequest method. To read the response, the client must call ReadResponseHeader and ReadResponseBody as a pair. Once the body has been read, it is the client's responsibility to call the Close method to close the connection. If a nil interface is passed to ReadResponseBody then the body of the response should be read and then discarded.

```
type ClientCodec interface {
// WriteRequest must be safe for concurrent use by multiple goroutines.
WriteRequest(*Request, interface{}) error
ReadResponseHeader(*Response) error
ReadResponseBody(interface{}) error
Close() error
}
```
The NewServerCodec method requires that we pass it a type that implements the ReadWriteCloser interface. As we do not have such a type passed to us as parameters in the httpHandler method we have defined our own type, HttpConn, which encapsulates the http.Request body, which implements io.Reader, and the ResponseWriter method, that implements io.Writer. We can then write our own methods that proxy the calls to the reader and writer creating a type that has the correct interface.

```go
type HttpConn struct {
	in  io.Reader
	out io.Writer
}

func (c *HttpConn) Read(req []byte) (n int, err error) {
	return c.in.Read(req)
}
func (c *HttpConn) Write(res []byte) (n int, err error) {
	return c.out.Write(res)
}
func (c *HttpConn) Close() error {
	return nil
}

```
### Server Handler
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
func PerformRequest() contract.HelloWorldResponse {
	r, _ := http.Post(
		"http://localhost:1234",
		"application/json",
		//where the magic happens
		bytes.NewBuffer([]byte(`{"id": 1, "method": "HelloWorldHandler.HelloWorld", "params": [{"name":"World"}]}`)),
	)
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	var response contract.HelloWorldResponse
	decoder.Decode(&response)

	return response
}
```

# Getting Started

```go
go run ./main.go
curl -X POST -H "Content-Type: application/json" -d '{"id": 1, "method": "HelloWorldHandler.HelloWorld", "params": [{"name":"World"}]}' http://localhost:1234
```

To run test
```
cd 09-PRC_HTTP_JSON
go test -v -run="none" -bench=. -benchtime="5s"
```

Output
```
{"id":1,"result":{"Message":"Hello World"},"error":null}
```