package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

// conn, err := net.Dial(connectionType string, address string)  //conn.close()
// reqLen, err := conn.Write([]byte("Request message to server"))
// resLen, err := conn.Read(received)  //received := make([]byte, 1024)

func main() {
	// Listen for incoming connections.
	conn, err := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT) // l as a Listener
	if err != nil {
		fmt.Println("Dial failed:", err.Error())
		os.Exit(1)
	}

	// Send the request --------------------------------
	_, err = conn.Write([]byte("This is a client message سلام")) //It's not unicode
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}

	// Receive the response --------------------------------
	// buffer to get data
	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		println("Read data failed:", err.Error())
		os.Exit(1)
	}

	println("Received server message:", string(received))

	conn.Close()
}
