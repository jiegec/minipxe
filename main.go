package main

import (
	"log"
	"net"

	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/insomniacslk/dhcp/dhcpv4/server4"
)

func handler(conn net.PacketConn, peer net.Addr, m *dhcpv4.DHCPv4) {
	log.Print(m.Summary())
}

func main() {
	addr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 53,
	}
	server, err := server4.NewServer("", addr, handler)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}

	log.Println("Starting dhcp server")
	server.Serve()
}
