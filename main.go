package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	dnsServer := serveDNS("127.0.0.1:53")
	tftpServer := serveTFTP("127.0.0.1:69")

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

}
