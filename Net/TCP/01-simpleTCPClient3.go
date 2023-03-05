package main

import (
	"context"
	"log"
	"net"
	"time"
)

const (
	HOST3 = "localhost"
	PORT3 = "3333"
	TYPE3 = "tcp"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	var d net.Dialer
	conn, err := d.DialContext(ctx, TYPE3, HOST3+":"+PORT3)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	if _, err := conn.Write([]byte("Hello, World!")); err != nil {
		log.Fatal(err)
	}
}
