package main

import (
	server "01-Introduction_to_Microservices/07-RPC/2-server"
	client "01-Introduction_to_Microservices/07-RPC/3-client"
	"fmt"
)

func main() {
	// Create a routine to start server.
	go server.StartServer()

	// Create multiple client to server instead of a client communicate to server.
	for i := 0; i < 10; i++ {
		// Right here, that's how it connects to the server.
		c := client.CreateClient()

		// Close connect to server
		//Possible resource leak, 'defer' is called in the 'for' loop then move it to end
		//defer c.Close()

		// Client make a request
		reply := client.PerformRequest(c)

		fmt.Println(reply.Message)

		// Close connect to server
		c.Close()
	}
}
