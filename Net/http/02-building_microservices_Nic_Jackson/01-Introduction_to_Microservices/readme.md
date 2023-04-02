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

## http.HandleFunc() 
for Request & Response
### json.Marshal() & json.Unmarshal()
```go
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    //---------Request
    body, err := io.ReadAll(r.Body)
    if err != nil {
      http.Error(w, "Bad request", http.StatusBadRequest)
      return
    }
    var request helloWorldRequest
    err = json.Unmarshal(body, &request)
    if err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }
    log.Println(request.Name)
	
    //---------Response
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
    encoder := json.NewEncoder(w)
    encoder.Encode(response)
}
```
## http.Handler
every struct implement this interface can use `func Handle(pattern string, handler Handler)`
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