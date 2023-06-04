package main

import (
	"fmt"
	"net/http"
	"time"
)

type PingInfo struct {
	url        string
	err        error
	latency    time.Duration
	statusCode int
}

func myPing(url string) PingInfo {
	start := time.Now()
	//we make a get request to the URL and calculate the time it took for the request to be completed
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err == nil {
		t := time.Since(start).Round(time.Millisecond)
		return PingInfo{url, nil, t, resp.StatusCode}
	}

	return PingInfo{url, err, 0, resp.StatusCode}
}

func main() {
	fmt.Println(myPing("https://www.google.com"))
}
