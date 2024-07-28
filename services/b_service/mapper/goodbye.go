package mapper

import (
	proto "benneighbour.com/proto/autogen/b_service/goodbye"
	"benneighbour.com/services/b_service/entity"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToEntity maps proto.GoodbyeReply to entity.Goodbye
func ToEntity(request *proto.GoodbyeRequest) *entity.Goodbye {
	return &entity.Goodbye{
		Message:   request.Message,
		Timestamp: timestamppb.Now().AsTime(),
		UserUuid:  uuid.UUID{},
	}
}

// ToReply maps entity.Goodbye to entity.Goodbye
func ToReply(entity *entity.Goodbye) *proto.GoodbyeReply {
	return &proto.GoodbyeReply{
		Message:   entity.Message,
		Timestamp: timestamppb.New(entity.Timestamp),
		UserUuid:  entity.UserUuid.String(),
	}
}
