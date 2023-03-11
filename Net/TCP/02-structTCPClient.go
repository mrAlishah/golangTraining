// client.go
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
)

type Message struct {
	ID   string
	Data string
}

func send(conn net.Conn) {
	// lets create the message we want to send accross
	msg := Message{ID: "Yo", Data: "Hello"}
	msgBuff := new(bytes.Buffer)

	// create a encoder object
	gobObj := gob.NewEncoder(msgBuff)
	// encode buffer and marshal it into a gob object
	gobObj.Encode(msg)

	conn.Write(msgBuff.Bytes())
}

func recv(conn net.Conn) {
	// create a temp buffer
	msg := make([]byte, 500)
	conn.Read(msg)

	// convert bytes into Buffer (which implements io.Reader/io.Writer)
	msgBuff := bytes.NewBuffer(msg)
	// creates a decoder object
	gobObj := gob.NewDecoder(msgBuff)

	tmpstruct := new(Message)
	// decodes buffer and unmarshals it into a Message struct
	gobObj.Decode(tmpstruct)

	fmt.Println(tmpstruct)
}

func main() {
	conn, _ := net.Dial("tcp", ":8081")

	// Uncomment to test timeout
	// time.Sleep(5 * time.Second)
	// return

	send(conn)
	recv(conn)
}
