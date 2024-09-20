package chat

import (
	"chat-server/internal/sys/validate"
	"chat-server/pkg/chat_v1"
	"context"
)

func (i *Implementation) Create(ctx context.Context, desc *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	err := validate.Validate(ctx, validateIds(desc.GetIds()))
	if err != nil {
		return nil, err
	}
	id, err := i.chatServ.Create(ctx, desc.Ids)
	if err != nil {
		return nil, err
	}
	return &chat_v1.CreateResponse{Id: id}, nil
}

func validateIds(ids []int64) validate.Condition {
	return func(ctx context.Context) error {
		for _, id := range ids {
			if id == 0 {
				return validate.NewValidationErrors("id must be grater then 0")
			}

		}
		return nil
	}
}
