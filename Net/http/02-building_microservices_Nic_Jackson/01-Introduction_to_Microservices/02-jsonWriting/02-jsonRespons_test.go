package main

import (
	"encoding/json"
	"fmt"
	"io"
	"testing"
)

type Response struct {
	Message string
}

// go test -bench=. ./02-jsonWriting
func BenchmarkHelloHandlerVariable(b *testing.B) {
	b.ResetTimer()

	var writer = io.Discard
	response := Response{Message: "Hello World"}

	for i := 0; i < b.N; i++ {
		data, _ := json.Marshal(response)
		fmt.Fprint(writer, string(data))
	}
}

func BenchmarkHelloHandlerEncoder(b *testing.B) {
	b.ResetTimer()

	var writer = io.Discard
	response := Response{Message: "Hello World"}

	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(writer)
		encoder.Encode(response)
	}
}

func BenchmarkHelloHandlerEncoderReference(b *testing.B) {
	b.ResetTimer()

	var writer = io.Discard
	response := Response{Message: "Hello World"}

	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(writer)
		encoder.Encode(&response)
	}
}
