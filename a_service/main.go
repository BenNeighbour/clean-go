package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"benneighbour.com/a_service/config"
	"benneighbour.com/a_service/transport"
)

// initializeTransport sets up and starts the gRPC server.
func initializeTransport(listener net.Listener) {
	server := grpc.NewServer()

	// Register the a_service transport with the gRPC server.
	transport.Register(server)

	// Enable server reflection for easier client interaction.
	reflection.Register(server)

	// Start serving gRPC on the specified listener.
	err := server.Serve(listener)

	// Check for error
	if err != nil {
		return
	}
}

func main() {
	// Parse command-line flags to initialize the environment.
	flag.Parse()

	// Listen on the specified port.
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *config.SERVER_PORT))

	// Check for error
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Start the transport.
	initializeTransport(listener)

	// Log the fact that we are indeed listening.
	log.Printf("server listening at %v", listener.Addr())
}
