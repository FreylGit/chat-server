package chat

import (
	"chat-server/internal/sys/validate"
	"chat-server/pkg/chat_v1"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, desc *chat_v1.DeleteRequest) (*emptypb.Empty, error) {
	err := validate.Validate(ctx, validateId(desc.GetId()))
	if err != nil {
		return nil, err
	}
	err = i.chatServ.Delete(ctx, desc.GetId())
	return &emptypb.Empty{}, err
}

func validateId(id int64) validate.Condition {
	return func(ctx context.Context) error {
		if id <= 0 {
			return validate.NewValidationErrors("id cannot be less than zero")
		}
		return nil
	}
}
