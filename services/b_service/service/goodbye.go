package service

import (
	proto "benneighbour.com/proto/autogen/b_service/goodbye"
	"benneighbour.com/services/b_service/mapper"
	"benneighbour.com/services/b_service/repository"
	"context"
)

// GoodbyeService defines the interface for the Goodbye service
type GoodbyeService interface {
	SayGoodbye(ctx context.Context, request *proto.GoodbyeRequest) (*proto.GoodbyeReply, error)
}

// GoodbyeServiceImpl is an implementation of the GoodbyeService interface
type GoodbyeServiceImpl struct {
	repo repository.GoodbyeRepository
}

// NewGoodbyeService creates a new instance of GoodbyeServiceImpl
func NewGoodbyeService(repo repository.GoodbyeRepository) GoodbyeService {
	return &GoodbyeServiceImpl{repo: repo}
}

// SayGoodbye handles the GoodbyeRequest and responds with a GoodbyeReply
func (s *GoodbyeServiceImpl) SayGoodbye(ctx context.Context, request *proto.GoodbyeRequest) (*proto.GoodbyeReply, error) {
	// Convert GoodbyeRequest to Goodbye entity
	entity, err := s.repo.Create(ctx, mapper.ToEntity(request))
	if err != nil {
		return nil, err
	}

	// Convert Goodbye entity back to gRPC reply and return
	return mapper.ToReply(entity), nil
}
