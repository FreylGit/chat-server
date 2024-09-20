package message

import (
	"chat-server/internal/model"
	"context"
)

func (s *serv) Create(ctx context.Context, message model.Message) error {
	err := s.messageRepo.Create(ctx, message)
	return err
}
