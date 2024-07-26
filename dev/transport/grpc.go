package transport

import (
	"benneighbour.com/dev/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func InitializeSuperTransport(listener net.Listener) {
	// Create a grpc server
	server := grpc.NewServer()

	// Register each other service's transport for the combined server
	for _, service := range config.Services {
		service.Register(server)
	}

	// Register the dev server for client reflection
	reflection.Register(server)

	// Serve grpc on our listener port
	err := server.Serve(listener)

	// Check for error
	if err != nil {
		return
	}
}
