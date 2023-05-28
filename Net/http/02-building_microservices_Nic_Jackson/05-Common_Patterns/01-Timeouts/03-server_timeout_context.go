// Reference: https://stackoverflow.com/questions/58736588/http-server-handlefunc-loop-on-timeout
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// main ctx timeout is earlier than worker time
// var WriteTimeout = 3 * time.Second

// main ctx timeout is longer than worker time > 6+2
var WriteTimeout = 20 * time.Second

// localhost:8080/timeout
// localhost:8080/home
func main() {

	router := http.NewServeMux()
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: WriteTimeout + 10*time.Millisecond, //10ms Redundant time
		IdleTimeout:  15 * time.Second,
	}

	router.HandleFunc("/slow", makeNormalRequest)
	router.HandleFunc("/timeout", makeTimeoutRequest)
	router.HandleFunc("/home", home)

	fmt.Printf("Server starting on port %v\n", server.Addr)
	server.ListenAndServe()
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("responding\n")
	ctx, _ := context.WithTimeout(context.Background(), WriteTimeout)
	worker, cancel := context.WithCancel(context.Background())

	var buffer string
	go func() {
		// do something
		time.Sleep(6 * time.Second)

		buffer = "ready all response\n"

		//do another
		time.Sleep(2 * time.Second)
		cancel()
		fmt.Printf("worker finish\n")
		//worker.Done()
	}()

	select {
	case <-ctx.Done():
		//add more friendly tips
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Timeout\n")
		return
	case <-worker.Done():
		w.Write([]byte(buffer))
		fmt.Fprint(w, "Worker done\n")
		fmt.Printf("writed\n")
		return
	}
}

func makeNormalRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "slow Funtion is starting\n")
	slowFunction()
}

func makeTimeoutRequest(w http.ResponseWriter, r *http.Request) {

	ctx, _ := context.WithTimeout(context.Background(), WriteTimeout)
	worker, cancel := context.WithCancel(context.Background())

	var buffer string
	go func() {
		slowFunction()
		cancel()
		fmt.Printf("worker finish\n")
	}()

	select {
	case <-ctx.Done():
		//add more friendly tips
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Timeout")
		fmt.Fprint(w, "Timeout\n")
		return
	case <-worker.Done():
		w.Write([]byte(buffer))
		fmt.Println("writed")
		fmt.Fprint(w, "Worker done\n")

		return
	}

}

func slowFunction() {
	for i := 0; i < 100; i++ {
		fmt.Println("Loop: ", i)
		time.Sleep(1 * time.Second)
	}
}
