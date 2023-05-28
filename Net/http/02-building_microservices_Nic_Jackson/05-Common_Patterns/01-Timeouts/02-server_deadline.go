package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/eapache/go-resiliency/deadline"
)

// localhost:8080/timeout
func main() {

	port := "8080"

	http.HandleFunc("/slow", makeNormalRequest)
	http.HandleFunc("/timeout", makeTimeoutRequest)

	fmt.Printf("Server starting on port %v\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	// Panic a entire server.
	if err != nil {
		panic(err.Error())
	}

}

func makeNormalRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "slow Funtion is starting\n")
	slowFunction()
}

func makeTimeoutRequest(w http.ResponseWriter, r *http.Request) {

	dl := deadline.New(3 * time.Second)
	err := dl.Run(func(stopper <-chan struct{}) error {
		slowFunction()
		return nil
	})

	switch err {
	case deadline.ErrTimedOut:
		fmt.Println("Timeout")
		fmt.Fprint(w, "Timeout\n")
		// To abort a request from your client, we must use panic & recovery.
		http.Error(w, "Timeout!", http.StatusBadRequest)
		panic(http.ErrAbortHandler) // terminate request (and any more handlers in the chain)
	default:
		fmt.Println(err)
		fmt.Fprint(w, err)
	}
}

func slowFunction() {
	for i := 0; i < 100; i++ {
		fmt.Println("Loop: ", i)
		time.Sleep(1 * time.Second)
	}
}
