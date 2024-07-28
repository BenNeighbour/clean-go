package main

import (
	"fmt"
	"log"
	"net"

	"benneighbour.com/dev/config"
	"benneighbour.com/dev/transport"
)

func main() {
	// Firstly, load into our environment by parsing the command-line flags
	config.Initialize()

	// Listen on the requested port
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *config.SERVER_PORT))

	// Check for error
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Start the "super" multi-service transport
	transport.InitializeSuperTransport(listener)

	// Log the fact that we are indeed listening
	log.Printf("Development server listening at %v", listener.Addr())
}
