package chat

import (
	"chat-server/internal/model"
	"context"
)

func (s *serv) SendMessage(ctx context.Context, message model.Message) error {
	err := s.messageServ.Create(ctx, message)
	if err != nil {
		return err
	}
	go s.chatStorage.NotifyMessage(message.ChatId, message)

	return nil
}
