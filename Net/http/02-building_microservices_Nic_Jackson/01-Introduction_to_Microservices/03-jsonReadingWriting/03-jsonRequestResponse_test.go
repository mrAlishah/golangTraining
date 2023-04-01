package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

// go test -bench=. ./03-jsonReadingWriting
func BenchmarkHelloWorld1Handler(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r, _ := http.Post(
			"http://localhost:8080/helloworld1",
			"application/json",
			bytes.NewBuffer([]byte(`{"Name":"World"}`)),
		)

		var response helloWorldResponse
		decoder := json.NewDecoder(r.Body)

		_ = decoder.Decode(&response)
	}
}

func BenchmarkHelloWorld2Handler(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r, _ := http.Post(
			"http://localhost:8080/helloworld2",
			"application/json",
			bytes.NewBuffer([]byte(`{"Name":"World"}`)),
		)

		var response helloWorldResponse
		decoder := json.NewDecoder(r.Body)

		_ = decoder.Decode(&response)
	}
}

func init() {
	go server()
}
