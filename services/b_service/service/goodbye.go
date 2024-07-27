package service

import (
	"context"

	proto "benneighbour.com/proto/autogen/b_service/goodbye"
)

// GoodbyeService defines the interface for the Goodbye service
type GoodbyeService interface {
	SayGoodbye(ctx context.Context, req *proto.GoodbyeRequest) (*proto.GoodbyeReply, error)
}
