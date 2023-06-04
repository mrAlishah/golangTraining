package handlers

import (
	"encoding/json"
	"fmt"
	"gopkg.in/alexcesaro/statsd.v2"
	"math/rand"
	"net/http"
	"stateDSample/requestToJSON"
	"time"
)

const (
	helloworldSuccess string = "kittenserver.helloworld.success"
	helloworldFailed  string = "kittenserver.helloworld.failed"
	helloworldTiming  string = "kittenserver.helloworld.timing"
)

// HelloWorldResponse defines a response returned from the /helloworld endpoint
type HelloWorldResponse struct {
	Message string `json:"message"`
}

// HelloWorldRequest defines a request sent to the /helloworld endpoint
type HelloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldHandler struct {
	statsd *statsd.Client
}

func NewHelloWorldHandler(statsd *statsd.Client) http.Handler {
	return &helloWorldHandler{statsd: statsd}
}

func (h *helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	timing := h.statsd.NewTiming()

	//get request data
	var request HelloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		// metric data
		h.statsd.Increment(helloworldFailed)

		// display request as a json for future logger
		message := requestToJSON.SerialzableRequest{r}
		fmt.Printf("handler: %s, status:%d, method: %s, request:%s", "Validation \n", http.StatusBadRequest, r.Method, message.ToJSON())

		// create bad response
		http.Error(rw, "Bad request", http.StatusBadRequest)

		return
	}

	name := request.Name
	fmt.Printf("request:%+v\n", request)
	response := HelloWorldResponse{Message: "Hello " + name}

	// create response
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)

	// metric data
	h.statsd.Increment(helloworldSuccess)

	// just simulate for delay
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)

	//metric data
	//h.statsd.NewTiming().Send(helloworldTiming)
	timing.Send(helloworldTiming)
}
