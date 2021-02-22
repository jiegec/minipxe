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

func serveDNS(address string) *server4.Server {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatal("Bad listen addr: ", err)
	}

	server, err := server4.NewServer("", addr, handler)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}

	log.Printf("Starting dhcp server at %s", address)
	go server.Serve()

	return server
}
