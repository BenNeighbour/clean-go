package transport

import (
	"context"
	"google.golang.org/grpc"
	"log"

	proto "benneighbour.com/proto/autogen/b_service/goodbye"
	"benneighbour.com/services/b_service/service"
)

// GoodbyeTransport implements the protobuf GoodbyeServer interface
type GoodbyeTransport struct {
	proto.UnimplementedGoodbyeServer
	service service.GoodbyeService
}

// NewGoodbyeTransport creates a new GoodbyeTransport instance with the given service
func NewGoodbyeTransport() *GoodbyeTransport {
	return &GoodbyeTransport{
		service: service.NewGoodbyeService(),
	}
}

// SayGoodbye handles the GoodbyeRequest and responds with a GoodbyeReply
func (s *GoodbyeTransport) SayGoodbye(ctx context.Context, req *proto.GoodbyeRequest) (*proto.GoodbyeReply, error) {
	// Delegate the request to the GoodbyeService
	reply, err := s.service.SayGoodbye(ctx, req)

	// Check for error
	if err != nil {
		log.Printf("Error handling GOODBYE request: %v", err)
		return nil, err
	}

	// Return the reply
	return reply, nil
}

func Register(server *grpc.Server) {
	// Register the GoodbyeTransport with its dependencies
	proto.RegisterGoodbyeServer(server, NewGoodbyeTransport())
}
