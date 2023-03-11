package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

/*
We will use ListenPacket() function to create a UDP server:
func ListenPacket(network, address string) (PacketConn, error): ListenPacket announces on the
local network address. The network must be "udp", "udp4", "udp6", "unixgram", or an IP transport.
The IP transports are "ip", "ip4", or "ip6" followed by a colon and a literal protocol number or a protocol name,
as in "ip:1" or "ip:icmp". For UDP and IP networks, if the host in the address parameter is empty or a literal unspecified IP address,
ListenPacket listens on all available IP addresses of the local system except multicast IP addresses.
*/
func main() {
	// listen to incoming udp packets
	udpServer, err := net.ListenPacket("udp", ":1053")
	if err != nil {
		log.Fatal(err)
	}
	defer udpServer.Close()
	fmt.Println("UDP Server started ...")

	for {
		buf := make([]byte, 1024)
		_, addr, err := udpServer.ReadFrom(buf)
		if err != nil {
			continue
		}
		go response(udpServer, addr, buf)
	}

}

func response(udpServer net.PacketConn, addr net.Addr, buf []byte) {
	time := time.Now().Format(time.ANSIC)
	responseStr := fmt.Sprintf("time received: %v. Your message: %v!", time, string(buf))

	udpServer.WriteTo([]byte(responseStr), addr)
}
