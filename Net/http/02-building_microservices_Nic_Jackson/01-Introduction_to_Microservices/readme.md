# http
## Create http.DefaultServeMux
`func HandleFunc(pattern string, handler func(ResponseWriter, *Request))`
The HandleFunc method creates a Handler type on the DefaultServeMux handler,mapping the path passed in the first parameter to the function in the second parameter.

`func ListenAndServe(addr string, handler Handler) error`
- Function: ListenAndServe takes two parameters, the TCP network address to bind the server to and the handler that will be used to route requests
- @Para1: port -- network address 8080 bind the server to all available IP addresses on port 8080.
- @Para2: nil  -- this is because we are using the DefaultServeMux handler, which we are setting up with our call to http.
- Since ListenAndServe blocks if the server starts correctly we will never exit on a successful start.
```go
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}
func main() {
    port := 8080
    http.HandleFunc("/helloworld", helloWorldHandler)
    log.Printf("Server starting on port %v\n", 8080)
	
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}	
```
Look at 01-basic_http_server sample. [here](./01-basic_http_server)
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

## 1) http.HandleFunc() 
Look at 02-jsonWriting sample. [here](./02-jsonWriting)
Look at 03-jsonReadingWriting sample. [here](./03-jsonReadingWriting)
for Request & Response
### json.Marshal() & json.Unmarshal()
```go
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    //---------Request 
    //Body request as a stream.
    //If we need the data contained in the body, we can simply read it into a byte array	
    body, err := io.ReadAll(r.Body)
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
    log.Println(request.Name)
	
    //---------Response
    // Take info request from client, and response to client.	
    response := helloWorldResponse{Message: "Hello " + request.Name}
    data, err := json.Marshal(response)
    if err != nil {
        panic("Ooops")
    }
    fmt.Fprint(w, string(data))
}
```

### json.NewEncoder() & json.NewDecoder()
This method is faster than previous method.
```go
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    //---------Request
    var request helloWorldRequest
    decoder := json.NewDecoder(r.Body)
	
    err := decoder.Decode(&request)
    if err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }
    log.Println(request.Name)
	
    //---------Response
    response := helloWorldResponse{Message: "Hello " + request.Name}
    // Create object encoder   
    encoder := json.NewEncoder(w)
    // Write JSON straight to an open writer	
    encoder.Encode(response)
}
```
## 2) http.Handler
every struct implement this interface can use `func Handle(pattern string, handler Handler)`
Look at 05-creatingHandlers sample. [here](./05-creatingHandlers)
```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

```go
//  Handler 1: Validation Handler. This is ServHTTP first chain
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
    // From h - validation handler, pass ResponseWriter, Request to Validation Handler.
    h.next.ServeHTTP(rw, r)
}
```

```go
// Handler 2: Reply message response. This is ServHTTP last chain
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
```

```go
//implement main()
func main(){
    handler := newValidationHandler(newHelloWorldHandler())

    // Note: we're going to build a handler, not function handler.
    http.Handle("/helloworld", handler)
    ...
}
```
## more Handler
### 3) FileServer
```go
func FileServer(root FileSystem) Handler
```
A FileServer function returns a handler that serves HTTP requests with the contents of the filesystem. This can be used to
serve static files such as images or other content that is stored on the file system.
```go
//Example
http.Handle("/images", http.FileServer(http.Dir("./images")))
```
Look at 04-staticFileHandler sample. [here](./04-staticFileHandler)
### 4) NotFoundHandler
```go
func NotFoundHandler() Handler
```
The NotFoundHandler function returns a simple request handler that replies to each request with a 404 page not found reply.

### 5) RedirectHandler
```go
func RedirectHandler(url string, code int) Handler
```
The RedirectHandler function returns a request handler that redirects each request it receives to the given URI using the given
status code. The provided code should be in the 3xx range and is usually StatusMovedPermanently, StatusFound, or
StatusSeeOther.

### 6) StripPrefix
```go
func StripPrefix(prefix string, h Handler) Handler
```
The StripPrefix function returns a handler that serves HTTP requests by removing the given prefix from the request URL's
path and then invoking h handler. If a path does not exist, then StripPrefix will reply with an HTTP 404 not found error.
```go
// To serve a directory on disk (/tmp) under an alternate URL 
//path (/tmpfiles/), use StripPrefix to modify the request 
//URL's path before the FileServer sees it: 
http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/tmp"))))
```

### 7) TimeoutHandler
```go
func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler
```
The TimeoutHandler function returns a Handler interface that runs h with the given time limit.
The new handler calls h.ServeHTTP to handle each request, but if a call runs for longer than its time limit, the handler responds
with a 503 Service Unavailable response with the given message (msg) in its body.

## Context
The problem with the previous pattern is that there is no way that you can pass the validated request from one handler to the
next without breaking the http.Handler interface,The Context type implements a safe method for accessing request-scoped data that is safe to use simultaneously by multiple Go routines.
### 1) Background
```go
func Background() Context
```
The Background method returns an empty context that has no values; it is typically used by the main function and as the top-
level Context.

### 2) WithCancel
```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
```
The WithCancel method returns a copy of the parent context with a cancel function, calling the cancel function releases
resources associated with the context and should be called as soon as operations running in the Context type are complete.

### 3) WithDeadline
```go
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
```
The WithDeadline method returns a copy of the parent context that expires after the current time is greater than deadline. At
this point, the context's Done channel is closed and the resources associated are released. It also passes back a CancelFunc
method that allows manual cancellation of the context.

### 4) WithTimeout
```go
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```
The WithTimeout method is similar to WithDeadline except you pass it a duration for which the Context type should exist.
Once this duration has elapsed, the Done channel is closed and the resources associated with the context are released.

### 5) WithValue
```go
func WithValue(parent Context, key interface{}, val interface{}) Context
```
The WithValue method returns a copy of the parent Context in which the val value is associated with the key. The Context
values are perfect to be used for request-scoped data.

### How to use context
Look at 06-useContext sample. [here](./06-useContext)
```go
type validationContextKey string //Look at sample 06

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
```

```go
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
```

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
  Until the current ponint, we don't know that what protocol does DefaultServeMux use?  
  However, we know DefaultServeMux use HTTP protocol.

When you use RPC standard:
* you can select your protocol such as: tcp, tcp4, tcp6, unix, or unixpacket.
* you also useing a given protocol and binding it to IP as same as DefaultServeMux

Look at 07-PRC sample. [here](./07-RPC)

## RPC over HTTP Transport protocol
As you know, 07-PRC is an example about communication between client and server without HTTP Protocol.
Right now, how we can implement application using HTTP by RPC.
In the instance that you need to use HTTP as your transport protocol then the rpc package can facilitate this by calling the HandleHTTP method.
Look at 08-PRC_HTTP sample. [here](./08-RPC_HTTP)

## JSON-RPC over HTTP 
Have you ever put a question that how we can communicate by JSON?
we will look at the package that provides a built-in codec for serializing and deserializing to the JSON-RPC standard. We will also look at how we can send these responses over HTTP.
Look at 09-PRC_HTTP_JSON sample. [here](./09-RPC_HTTP_JSON)