package transport

import (
	"context"
	"log"

	"google.golang.org/grpc"

	proto "benneighbour.com/proto/b_service"
)

// GoodbyeTransport implements the protobuf GoodbyeServer interface
type GoodbyeTransport struct {
	proto.UnimplementedGoodbyeServer
}

// SayGoodbye handles the GoodbyeRequest and responds with a GoodbyeReply
func (s *GoodbyeTransport) SayGoodbye(ctx context.Context, req *proto.GoodbyeRequest) (*proto.GoodbyeReply, error) {
	// Extract the name from the request
	name := req.GetName()

	// TODO - This belongs in the service
	log.Printf("Received GOODBYE request with name: %v", name)

	// TODO - This belongs in the service
	reply := &proto.GoodbyeReply{
		Message: "Goodbye " + name,
	}

	// Return the reply
	return reply, nil
}

// Register registers the GoodbyeTransport with the provided gRPC server
func Register(server *grpc.Server) {
	proto.RegisterGoodbyeServer(server, &GoodbyeTransport{})
}
