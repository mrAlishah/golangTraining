package main

import (
	server "01-Introduction_to_Microservices/09-RPC_HTTP_JSON/2-server"
	client "01-Introduction_to_Microservices/09-RPC_HTTP_JSON/3-client"
	"testing"
)

// cd ./01-Introduction_to_Microservices/09-RPC_HTTP_JSON
// go test -v -run="none" -bench=. -benchtime="5s"
func BenchmarkHelloWorldHandlerJSONRPC(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = client.PerformRequest()

	}
}

func init() {
	// start the server
	go server.StartServer()
}
