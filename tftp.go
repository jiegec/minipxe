package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/pin/tftp"
)

// readHandler is called when client starts file download from server
func readHandler(filename string, rf io.ReaderFrom) error {
	log.Printf("Client reading file %s", filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Got error when opening %s: %s", filename, err)
		return err
	}

	n, err := rf.ReadFrom(file)
	if err != nil {
		log.Printf("Got error when reading %s: %s", filename, err)
		return err
	}

	log.Printf("%d bytes sent", n)
	return nil
}

func serveTFTP(address string) *tftp.Server {
	s := tftp.NewServer(readHandler, nil)
	s.SetTimeout(5 * time.Second)

	go func() {
		log.Printf("Starting tftp server at %s", address)
		err := s.ListenAndServe(address)
		if err != nil {
			log.Fatal("Failed to start tftp server: ", err)
		}
	}()

	return s
}
