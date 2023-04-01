# http
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

# Request & Response
## json.Marshal() & json.Unmarshal()
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

## json.NewEncoder() & json.NewDecoder()
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