package main

import (
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

/*
func ResolveUDPAddr(network, address string) (*UDPAddr, error): ResolveUDPAddr returns an address of UDP end point. The network must be a UDP network name.
func DialUDP(network string, laddr, raddr *UDPAddr) (*UDPConn, error): DialUDP acts like Dial for UDP networks. Dial connects to the address on the named network.
*/
func main() {
	udpServer, err := net.ResolveUDPAddr("udp", ":1053")

	if err != nil {
		println("ResolveUDPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, udpServer)
	if err != nil {
		println("Listen failed:", err.Error())
		os.Exit(1)
	}

	//close the connection
	defer conn.Close()

	_, err = conn.Write([]byte("This is a UDP message"))
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}

	// buffer to get data
	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		println("Read data failed:", err.Error())
		os.Exit(1)
	}

	println(string(received))
}
