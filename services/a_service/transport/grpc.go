package transport

import (
	"context"
	"log"

	"google.golang.org/grpc"

	proto "benneighbour.com/proto/a_service"
)

// HelloTransport implements the protobuf HelloServer interface
type HelloTransport struct {
	proto.UnimplementedHelloServer
}

// SayHello handles the HelloRequest and responds with a HelloReply
func (s *HelloTransport) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	// Extract the name from the request
	name := req.GetName()

	// TODO - This belongs in the service
	log.Printf("Received HELLO request with name: %v", name)

	// TODO - This belongs in the service
	reply := &proto.HelloReply{
		Message: "Hello " + name,
	}

	// Return the reply
	return reply, nil
}

// Register registers the HelloTransport with the provided gRPC server
func Register(server *grpc.Server) {
	proto.RegisterHelloServer(server, &HelloTransport{})
}
