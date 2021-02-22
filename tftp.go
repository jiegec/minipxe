package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/pin/tftp"
)

// readHandler is called when client starts file download from server
func readHandler(filename string, rf io.ReaderFrom) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	n, err := rf.ReadFrom(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	fmt.Printf("%d bytes sent\n", n)
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
