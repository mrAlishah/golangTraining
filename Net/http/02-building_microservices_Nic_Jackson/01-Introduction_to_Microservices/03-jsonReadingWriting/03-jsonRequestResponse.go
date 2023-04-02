package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type helloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

/*************************************************************************************************/
/* http.HandleFunc:
/* 1) body, err := io.ReadAll(r.Body)
/* 2) json.Unmarshal(body, &request)       //var request helloWorldRequest
/* 3) data, err := json.Marshal(response)  //response := helloWorldResponse{Message: "Hello " + request.Name}
/*************************************************************************************************/
// This method is slower than json.Decoder/Encoder
func helloWorldUnmarshalHandler(w http.ResponseWriter, r *http.Request) {

	//The JSON that has been sent with the request is accessible in the Body field. Body implements the interface io.ReadCloser as a stream and does
	//not return a []byte or a string. If we need the data contained in the body, we can simply read it into a byte array,
	//Here is something we'll need to remember. We are not calling Body.Close(), if we were making a call with a client we would need to do this as it is
	//not automatically closed, however, when used in a ServeHTTP handler, the server automatically closes the request stream.

	//---------Request
	// Body request as a stream.
	//  If we need the data contained in the body, we can simply read it into a byte array
	// body, err := ioutil.ReadAll(r.Body)  Go.1.15 and earlier
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
	//log.Println(request.Name)

	//---------Response
	// Take info request from client, and response to client.
	response := helloWorldResponse{Message: "Hello " + request.Name}

	// Create object encoder
	encoder := json.NewEncoder(w)
	// Write JSON straight to an open writer
	encoder.Encode(response)
	/*Or
	  data, err := json.Marshal(response)
	  if err != nil {
	      panic("Ooops")
	  }

	  fmt.Fprint(w, string(data))
	*/
}

/*************************************************************************************************/
/* http.HandleFunc:
/* 1) decoder := json.NewDecoder(r.Body) // Reading
/* 2) err := decoder.Decode(&request)    // var request helloWorldRequest
/* 3) encoder := json.NewEncoder(w)      // Writing
/* 4) encoder.Encode(response)           // response := helloWorldResponse{Message: "Hello " + request.Name}
/*************************************************************************************************/
// This method is fast
func helloWorldDecoderHandler(w http.ResponseWriter, r *http.Request) {
	//---------Request
	var request helloWorldRequest
	// Like Marshall, we can use NewDecoder() func to create object for decoding.
	// Create object decoder to decode Body request.
	decoder := json.NewDecoder(r.Body)

	//  Decode Body request and put to request.
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	//log.Println(request.Name)

	//---------Response
	// After that, we can use it directly.
	response := helloWorldResponse{Message: "Hello " + request.Name}

	// Create object encoder.
	encoder := json.NewEncoder(w)

	// Write JSON directly to response.
	encoder.Encode(response)
}

// curl localhost:8080/helloworld1 -d '{"name":"Alishah"}'
// curl -X POST localhost:8080/helloworld1 -H "Content-Type: application/json" -d '{"name":"Alishah"}'
func main() {
	server()
}

func server() {
	port := 8080

	http.HandleFunc("/helloworld1", helloWorldUnmarshalHandler)
	//json.NewEncoder(w) is Faster than json.Marshal(response)
	//json.NewDecoder(r.Body) is Faster than json.Unmarshal(body, &request)
	http.HandleFunc("/helloworld2", helloWorldDecoderHandler)

	log.Printf("Server starting on port %v\n", 8080)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
