package service

import (
	"context"
	"fmt"
	"log"

	proto "benneighbour.com/proto/autogen/b_service/goodbye"
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
	// Extract the name from the request and log it
	name := req.GetName()
	log.Printf("Received GOODBYE request with name: %s", name)

	// Return the response with a formatted message
	return &proto.GoodbyeReply{
		Message: fmt.Sprintf("Goodbye %s", name),
	}, nil
}
