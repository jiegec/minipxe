package main

import (
	"log"
	"net"
	"time"

	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/insomniacslk/dhcp/dhcpv4/server4"
)

var gatewayIP net.IP
var serverIP net.IP
var clientIP net.IP
var subnetMask net.IPMask
var tftpBoot string

func handler(conn net.PacketConn, peer net.Addr, m *dhcpv4.DHCPv4) {
	log.Printf("Received packet local %s peer %s: %s", conn.LocalAddr(), peer, m.Summary())
	if m.MessageType() == dhcpv4.MessageTypeDiscover {
		reply, err := dhcpv4.New(
			dhcpv4.WithReply(m),
			dhcpv4.WithGatewayIP(gatewayIP),
			dhcpv4.WithServerIP(serverIP),
			dhcpv4.WithYourIP(clientIP),
			dhcpv4.WithMessageType(dhcpv4.MessageTypeOffer),
			dhcpv4.WithOptionCopied(m, dhcpv4.OptionClientIdentifier),
			dhcpv4.WithOption(dhcpv4.OptSubnetMask(subnetMask)),
			dhcpv4.WithOption(dhcpv4.OptIPAddressLeaseTime(time.Duration(24*time.Hour))),
			dhcpv4.WithOption(dhcpv4.OptRouter(gatewayIP)),
		)
		if err != nil {
			log.Print("Got error when constructing reply: ", err)
			return
		}
		log.Print("Sent: ", reply.Summary())
		conn.WriteTo(reply.ToBytes(), peer)
	} else if m.MessageType() == dhcpv4.MessageTypeRequest {
		reply, err := dhcpv4.New(
			dhcpv4.WithReply(m),
			dhcpv4.WithGatewayIP(gatewayIP),
			dhcpv4.WithServerIP(serverIP),
			dhcpv4.WithYourIP(clientIP),
			dhcpv4.WithMessageType(dhcpv4.MessageTypeAck),
			dhcpv4.WithOptionCopied(m, dhcpv4.OptionClientIdentifier),
			dhcpv4.WithOption(dhcpv4.OptSubnetMask(subnetMask)),
			dhcpv4.WithOption(dhcpv4.OptRouter(gatewayIP)),
			dhcpv4.WithOption(dhcpv4.OptIPAddressLeaseTime(time.Duration(24*time.Hour))),
			dhcpv4.WithOption(dhcpv4.OptTFTPServerName(serverIP.String())),
			dhcpv4.WithOption(dhcpv4.OptBootFileName(tftpBoot)),
		)
		if err != nil {
			log.Print("Got error when constructing reply: ", err)
			return
		}
		log.Print("Sent: ", reply.Summary())
		conn.WriteTo(reply.ToBytes(), peer)
	}
}

func serveDHCP(networkInterface string, address string) *server4.Server {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatal("Bad listen addr: ", err)
	}

	server, err := server4.NewServer(networkInterface, addr, handler)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}

	log.Printf("Starting dhcp server at %s", address)
	go server.Serve()

	return server
}
