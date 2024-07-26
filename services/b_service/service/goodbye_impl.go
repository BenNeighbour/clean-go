package service

import (
	"context"
	"fmt"
	"log"

	proto "benneighbour.com/proto/b_service"
)

// GoodbyeServiceImpl is an implementation of the GoodbyeService interface
type GoodbyeServiceImpl struct {
	// Dependencies (e.g., repositories, clients)
}

// NewGoodbyeService creates a new instance of GoodbyeServiceImpl
func NewGoodbyeService() GoodbyeService {
	return &GoodbyeServiceImpl{}
}

// SayGoodbye handles the GoodbyeRequest and responds with a GoodbyeReply
func (s *GoodbyeServiceImpl) SayGoodbye(ctx context.Context, req *proto.GoodbyeRequest) (*proto.GoodbyeReply, error) {
	// Extract the name from the request
	name := req.GetName()

	// Structured logging using a logging library
	log.Printf("Received GOODBYE request with name: %s", name)

	// Business logic: generate the response message
	message := fmt.Sprintf("Goodbye %s", name)

	// Create the response
	reply := &proto.GoodbyeReply{
		Message: message,
	}

	// Return the response
	return reply, nil
}
