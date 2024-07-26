package service

import (
	"context"

	proto "benneighbour.com/proto/b_service"
)

// GoodbyeService defines the interface for the Goodbye service
type GoodbyeService interface {
	SayGoodbye(ctx context.Context, req *proto.GoodbyeRequest) (*proto.GoodbyeReply, error)
}
