package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/urfave/cli/v2"
)

func action(c *cli.Context) error {
	dnsServer := serveDNS(c.String("dns-listen"))
	tftpServer := serveTFTP(c.String("tftp-listen"))

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	exit := false
	for !exit {
		select {
		case sig := <-signalChannel:
			log.Printf("Received signal %s, quitting", sig)
			dnsServer.Close()
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
				Name:  "dns-listen",
				Value: ":53",
				Usage: "Listen address for DNS server",
			},
			&cli.StringFlag{
				Name:  "tftp-listen",
				Value: ":69",
				Usage: "Listen address for TFTP server",
			},
		},
		Action: action,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
