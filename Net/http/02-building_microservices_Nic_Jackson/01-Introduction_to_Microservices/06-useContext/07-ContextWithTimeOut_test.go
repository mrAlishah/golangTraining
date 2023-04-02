package main

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

/*
************************************************************************************************
Server manages the lifecycle of the context automatically cancelling it when the client
connection closes. For outbound requests, Context controls cancellation, by this we mean that if
we cancel the Context() method we can cancel the outgoing request

func fetchGoogle(): find request to cancel that connection. This is a function running in background

************************************************************************************************
*/

func TestFetchGoogle(t *testing.T) {
	r, _ := http.NewRequest("GET", "https://google.com", nil)

	timeoutRequest, cancelFunc := context.WithTimeout(r.Context(), 1*time.Millisecond)
	defer cancelFunc()

	r = r.WithContext(timeoutRequest)

	_, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
