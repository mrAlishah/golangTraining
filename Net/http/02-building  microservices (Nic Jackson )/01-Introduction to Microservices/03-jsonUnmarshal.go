package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type helloWorld03Request struct {
	Name string `json:"name"`
}

type helloWorld03Response struct {
	Message string `json:"message"`
}

func helloWorld03UnmarshalHandler(w http.ResponseWriter, r *http.Request) {

	//The JSON that has been sent with the request is accessible in the Body field. Body implements the interface io.ReadCloser as a stream and does
	//not return a []byte or a string. If we need the data contained in the body, we can simply read it into a byte array,
	//Here is something we'll need to remember. We are not calling Body.Close(), if we were making a call with a client we would need to do this as it is
	//not automatically closed, however, when used in a ServeHTTP handler, the server automatically closes the request stream.

	//---------Request
	// body, err := ioutil.ReadAll(r.Body)  Go.1.15 and earlier
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var request helloWorld03Request
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	log.Println(request.Name)

	//---------Response
	response := helloWorld03Response{Message: "Hello " + request.Name}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
	/*Or
	  data, err := json.Marshal(response)
	  if err != nil {
	      panic("Ooops")
	  }

	  fmt.Fprint(w, string(data))
	*/
}

func helloWorld03DecoderHandler(w http.ResponseWriter, r *http.Request) {
	//---------Request
	var request helloWorld03Request
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	log.Println(request.Name)

	//---------Response
	response := helloWorld03Response{Message: "Hello " + request.Name}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

// curl localhost:8080/helloworld1 -d '{"name":"Alishah"}'
// curl -X POST localhost:8080/helloworld1 -H "Content-Type: application/json" -d '{"name":"Alishah"}'
func main() {
	port := 8080

	http.HandleFunc("/helloworld1", helloWorld03UnmarshalHandler)
	//json.NewEncoder(w) is Faster than json.Marshal(response)
	//json.NewDecoder(r.Body) is Faster than json.Unmarshal(body, &request)
	http.HandleFunc("/helloworld2", helloWorld03DecoderHandler)

	log.Printf("Server starting on port %v\n", 8080)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
