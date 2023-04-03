# Introduction
This chapter 1 will help you understand some lesson below.  

In this example, we are going to create an HTTP server with a single endpoint that returns static text represented by the JSON  
standard, this will introduce the basic functions of the HTTP server and handlers.  
We will then modify this endpoint to accept a request that is encoded in JSON and using the encoding/json package return a response to the client.  
We will also examine how the routing works by adding a second endpoint that returns a simple image.  
By the end of this chapter, you will have a fundamental grasp of the basic packages and how you can use them to quickly and
efficiently build a simple microservice.
## Build-web-server

To run this example:  
```
$go run ./folderName/fileName.go
```

You should now see the application output:  
```
2022/02/08 23:08:54 Server starting on port 8080
```

To check this instance running on process.  
```
$ps -aux | grep 'go run
```
## Json
Thanks to the encoding /json [package](https://pkg.go.dev/encoding/json), which is built into the standard library encoding and decoding JSON to and from Go types is both fast and easy.
It implements the simplistic Marshal and Unmarshal functions; however, if we need them, the package also provides Encoder and Decoder types that allow us greater control when reading and writing

### Marshalling-Go-structs-to-JSON 
Json in Golang is powerful, we need to remember some feature for coding:
* Change the output field to be "message"
```
    Message string `json:"message"`
```
* Do not use as an output this field in struct.
```
    Author string `json:"-"`
```
* Do not output the field if the value is empty.
```
    Date string `json:",omitempty"`
```
* convert output to a string and rename "id".
```
    Id int `json:"id, string"`
```
#### Using Marshalling

```go
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	// Create message response following member in struct.
	response := helloWorldResponse{Message: "HelloWorld"}

	// Convert to json by using Marshal
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops")
	}
	// Write(p []byte) (n int, err error) only accept bytes
	// So we use Fprint to convert json to bytes.
	fmt.Fprint(w, string(data))
}
```

To run it
```
Open brower:
> http://localhost:8080/helloworld

Or open CMD:
curl http://localhost:8080/helloworld
```

Output
```
{"Message":"Hello World"}
```

#### Encoder for Marshalling
If you familiar with Multiple Marshalling, maybe you ask the developer that:
> Is there any better way to send our data to the output stream without marshalling to a temporary object before we return it? 

The answer is yes.
The encoding/json package has a function called NewEncoder this returns us an Encoder object that can be used to write JSON straight to an open writer and guess what?  
Using Encode  
```go
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Create message response following member in struct.
	response := helloWorldResponse{Message: "HelloWorld"}
	// Create object encoder.
	encoder := json.NewEncoder(w)
	// Write JSON straight to an open writer
	encoder.Encode(response)
}
```
### Unmarshalling-JSON-to-Go-structs
```go
	func Unmarshal(data []byte, v interface{}) error
```
This function will allocate maps, slices, and pointers as required  
#### HTML-request-format
* io.ReadCloser as a stream and does not return a []byte or a string
```go
type Requests struct {
    ...
    // Method specifies the HTTP method (GET, POST, PUT, etc.).
    Method string
    // Header contains the request header fields received by the server. The type Header is a link to map[string] []string.
    Header Header
    // Body is the request's body.
    Body io.ReadCloser
    ...
}
```

#### Unmarshalling
```go
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	// Body request as a stream.
	//  If we need the data contained in the body, we can simply read it into a byte array
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Convert json body request to local request.
	var request helloWorldRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	// Take info request from client, and response to client.
	response := helloWorldResponse{Message: "Hello " + request.Name}

	// Create object encoder
	encoder := json.NewEncoder(w)
	// Write JSON straight to an open writer
	encoder.Encode(response)
}
```
To run it
```
Open brower:
> http://localhost:8080/helloworld <=== Need to fix

Or open CMD:
curl -id "{\"name\":\"Thong\"}" 127.0.0.1:8080/helloworld
```

Output
```
HTTP/1.1 200 OK
Date: Sun, 13 Mar 2022 15:04:15 GMT
Content-Length: 26
Content-Type: text/plain; charset=utf-8

{"message":"Hello Thong"}
```

#### Decoder for UnMarshalling
As the same as Marshalling, this also have a decoder object to decode all request from client.
```go
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "Hello " + request.Name}

	encoder := json.NewEncoder(w)

	encoder.Encode(response)
}
```

## Routing
Even a simple microservice will need the capability to route requests to different handlers dependent on the requested path or
method.
- In Go this is handled by the DefaultServeMux method which is an instance of ServerMux.
- When we call the http.HandleFunc("/helloworld", helloWorldHandler) package function we are actually just indirectly calling http.DefaultServerMux.HandleFunc(…).

There are two functions to adding handlers to a ServerMux handler:
- Function handler
> func HandlerFunc(pattern string, handler func(ResponseWriter, *Request))
- Handler
> func Handle(pattern string, handler Handler)

## RPC
Remote Procedure Call(RPC) in Operating System is a powerful technique for constructing distributed, client-server based applications.  
More details: [here](https://www.geeksforgeeks.org/remote-procedure-call-rpc-in-operating-system/)  

### Protocol
Firstly, we need to know about architecture of protocol: [here](#https://www.digitalocean.com/community/tutorials/http-1-1-vs-http-2-what-s-the-difference)  
And you know that the data will transfer to internet through 4 layer:  
* Application Layer (HTTP)
* Transport Layer (TCP)
* Network Layer (IP)
* Data Link Layer
Until the current pint, we don't know that what protocol does DefaultServeMux use?  
However, we know DefaultServeMux use HTTP protocol.  
  
When you use RPC standard:
* you can select your protocol such as: tcp, tcp4, tcp6, unix, or unixpacket.
* you also use a given protocol and binding it to IP as same as DefaultServeMux 

### Simple RPC example
#### Server
To register a handler into server in RPC API.
```go
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld)
```

To make server listen client.
```go
	l, err := net.Listen("("tcp",", fmt.Sprintf(":%(":%v",", port))
```

To accept a connection between client and server, and block to wait client complete
```go
	for {
		conn, _ := l.Accept()
			go rpc.ServeConn(conn)
		}
	}
```

#### Client
How client can connect to server without HTTP protocol?
```go
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))
```

to make a request from Client
```go
	err := client.Call("HelloWorldHandler.HelloWorld", args, &reply)
```

### RPC over HTTP
As you know, Simple-RPC-example is an example about communication between client and server without HTTP Protocol.
Right now, how we can implement application using HTTP by RPC.

#### Server
To make RPC over HTTP
```go
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld)
	rpc.HandleHTTP()
```

And then, we need HTTP serve our server
```go
	l, err := net.Listen("("tcp",", fmt.Sprintf(":%(":%v",", port))
	http.Serve(l, nil)
```

#### Client
To make connection through HTTP.
```go
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("localhost:%v", port))
```

### JSON-RPC over HTTP
Have you ever put a question that how we can communicate by JSON?  

#### Server
To create json on handler
```go
	serverCodec := jsonrpc.NewServerCodec(&HttpConn{in: r.Body, out: w})
```

#### Client
To create a request by json
```go
	r, _ := http.Post(
		"http://localhost:1234",
		"application/json",
		bytes.NewBuffer([]byte(`{"id": 1, "method": "HelloWorldHandler.HelloWorld", "params": [{"name":"World"}]}`)),
	)
```
