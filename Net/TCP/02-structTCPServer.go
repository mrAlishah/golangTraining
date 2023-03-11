// server.go
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// Create your custom data struct
type Message struct {
	ID   string
	Data string
}

func logerr(err error) bool {
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			log.Println("read timeout:", err)
		} else if err == io.EOF {
		} else {
			log.Println("read error:", err)
		}
		return true
	}
	return false
}

// []byte -> *byte.Buffer -> *Decoder (gobObj) -> Decode(Struct)
// 1) msg := make([]byte, 500) -> conn.Read(msg) -> msgBuff := bytes.NewBuffer(msg) -> gobObj := gob.NewDecoder(msgBuff)
// 2) tmpstruct := new(Message)-> gobObj.Decode(tmpstruct)
func read(conn net.Conn) {
	// create a temp buffer
	msg := make([]byte, 500)

	// loop through the connection to read incoming connections. If you're doing by
	// directional, you might want to make this into a seperate go routine
	for {
		_, err := conn.Read(msg)
		if logerr(err) {
			break
		}

		// convert bytes into Buffer (which implements io.Reader/io.Writer)
		msgBuff := bytes.NewBuffer(msg)
		// creates a decoder object
		gobObj := gob.NewDecoder(msgBuff)

		tmpstruct := new(Message)
		// decodes buffer and unmarshals it into a Message struct
		gobObj.Decode(tmpstruct)

		// lets print out!
		fmt.Println(tmpstruct)
		return
	}
}

// *byte.Buffer -> *Encoder (gobObj) -> Encode(Struct)   "(msg~=*Buffer)"
// 1) msgBuff := new(bytes.Buffer) -> gob.NewEncoder(msgBuff) -> gobObj.Encode(msg) ->
// 2) conn.Write(msgBuff.Bytes())
func resp(conn net.Conn) {
	msg := Message{ID: "Yo", Data: "Hello back"}
	msgBuff := new(bytes.Buffer)

	// create a encoder object
	gobObj := gob.NewEncoder(msgBuff)
	// encode buffer and marshal it into a gob object
	gobObj.Encode(msg)

	conn.Write(msgBuff.Bytes())
	conn.Close()
}

func handle(conn net.Conn) {
	timeoutDuration := 2 * time.Second
	fmt.Println("Launching server...")
	conn.SetReadDeadline(time.Now().Add(timeoutDuration))

	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)

	read(conn)
	resp(conn)
}

func main() {
	server, _ := net.Listen("tcp", ":8081")
	fmt.Println("Starting server...")
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println("Connection error: ", err)
			return
		}
		go handle(conn)
	}
}
