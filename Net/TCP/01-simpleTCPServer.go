package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	SRV_HOST = "localhost"
	SRV_PORT = "3333"
	SRV_TYPE = "tcp"
)

// listener err := net.Listen(connectionType string, address string)  //listener.close()
// conn, err:= listener.Accept() //conn.close()
// reqLen, err := conn.Read(buf)  //buf := make([]byte, 1024)
// resLen, err :=conn.Write([]byte("Message to client for client request."))

// use command "echo -n "test out the server" | nc localhost 3333" for client
func main() {
	// Listen for incoming connections.
	l, err := net.Listen(SRV_TYPE, SRV_HOST+":"+SRV_PORT) // l as a Listener
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closSes.
	defer l.Close()

	fmt.Println("Listening on " + SRV_HOST + ":" + SRV_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept() // conn as a connection
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Read the request --------------------------------
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)

	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	log.Printf("server IP: %s , Client/Reciever IP: %s", conn.LocalAddr().String(), conn.RemoteAddr().String())
	log.Printf("Message received:%d , %s\n", reqLen, string(buf[:reqLen]))

	// Send the response --------------------------------
	// Send a response back to person contacting us.
	resLen, err := conn.Write([]byte("Message received.\n"))
	if err != nil {
		fmt.Println("Error writing:", err.Error())
	}
	fmt.Printf("Message sent:%d \n", resLen)

	// Close the connection when you're done with it.
	conn.Close()
}
