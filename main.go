package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/urfave/cli/v2"
)

func action(c *cli.Context) error {
	dhcpServer := serveDHCP(c.String("dhcp-interface"), c.String("dhcp-listen"))
	tftpServer := serveTFTP(c.String("tftp-listen"))
	gatewayIP = net.ParseIP(c.String("dhcp-gateway-ip"))
	serverIP = net.ParseIP(c.String("dhcp-server-ip"))
	var clientCIDR *net.IPNet
	var err error
	clientIP, clientCIDR, err = net.ParseCIDR(c.String("dhcp-client-cidr"))
	if err != nil {
		log.Fatal("Bad client cidr: ", err)
	}
	subnetMask = clientCIDR.Mask
	tftpRoot = c.String("tftp-root")
	tftpBoot = c.String("tftp-boot")
	ipxeConfig = c.String("ipxe-config")

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	exit := false
	for !exit {
		select {
		case sig := <-signalChannel:
			log.Printf("Received signal %s, quitting", sig)
			dhcpServer.Close()
			tftpServer.Shutdown()
			exit = true
		}
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:  "minipxe",
		Usage: "A minimal DNS and TFTP server for PXE",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "dhcp-listen",
				Value: ":67",
				Usage: "Listen address for DNS server",
			},
			&cli.StringFlag{
				Name:  "dhcp-server-ip",
				Value: "192.168.0.1",
				Usage: "Announced server ip address",
			},
			&cli.StringFlag{
				Name:  "dhcp-gateway-ip",
				Value: "192.168.0.1",
				Usage: "Announced gateway ip address",
			},
			&cli.StringFlag{
				Name:  "dhcp-client-cidr",
				Value: "192.168.0.100/24",
				Usage: "IP client cidr (only one is supported)",
			},
			&cli.StringFlag{
				Name:  "dhcp-interface",
				Value: "minipxe-test",
				Usage: "Network interface",
			},
			&cli.StringFlag{
				Name:  "tftp-listen",
				Value: ":69",
				Usage: "Listen address for TFTP server",
			},
			&cli.StringFlag{
				Name:  "tftp-root",
				Value: "root-ipxe",
				Usage: "Directory root to serve TFTP contents",
			},
			&cli.StringFlag{
				Name:  "tftp-boot",
				Value: "ipxe.efi",
				Usage: "TFTP boot file",
			},
			&cli.StringFlag{
				Name:  "ipxe-config",
				Value: "tftp://192.168.0.1/ipxe.cfg",
				Usage: "iPXE config path",
			},
		},
		Action: action,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
