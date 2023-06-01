package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/VividCortex/ewma"
)

var ma ewma.MovingAverage

// Step 1: Define threshold for calling API
// var threshold = 1000 * time.Millisecond  //corrected healthy
var threshold = 1000 * time.Nanosecond //invoke unhealthy

var timeout = 1000 * time.Millisecond
var resetting = false
var resetMutex = sync.RWMutex{}

func main() {
	ma = ewma.NewMovingAverage()

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/health", healthHandler)

	http.ListenAndServe(":8080", nil)
}

func mainHandler(rw http.ResponseWriter, r *http.Request) {
	// Step 2: Start time when call this API
	startTime := time.Now()

	// Step 3: Call health check API to check status
	// if the current moving average is greater than a defined threshold if the service is not healthy we return the status code
	// StatusServiceUnavailableS
	if !isHealthy() {
		respondServiceUnhealthy(rw)
		return
	}

	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "Average request time: %f (ms)\n", ma.Value()/1000000)

	duration := time.Now().Sub(startTime)
	ma.Add(float64(duration))
}

func healthHandler(rw http.ResponseWriter, r *http.Request) {
	if !isHealthy() {
		rw.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	fmt.Fprint(rw, "OK")
}

func isHealthy() bool {
	return (ma.Value() < float64(threshold))
}

// This func() will return response service for unhealthy
func respondServiceUnhealthy(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusServiceUnavailable)
	fmt.Println("service is unhealthy we need to sleep to give the service time to recover and then reset the average")

	// Step 4: call RLock() to obtain a lock on the resetMutex
	// we need this lock as when the service is unhealthy we need to sleep to
	// give the service time to recover and then reset the average
	resetMutex.RLock()
	defer resetMutex.RUnlock()

	// Step 5: resetting is false, call sleepAndResetAverage()
	if !resetting {
		go sleepAndResetAverage()
	}
}

func sleepAndResetAverage() {
	// Step 6: lock() this process, set resetting is true, and unlock()
	resetMutex.Lock()
	resetting = true
	resetMutex.Unlock()

	// Step 7: Sleep until timeout
	time.Sleep(timeout)
	// Step 8: Create a new moving average/
	ma = ewma.NewMovingAverage()

	// Step 9: Set variable as same as the first time
	resetMutex.Lock()
	resetting = false
	resetMutex.Unlock()
}
