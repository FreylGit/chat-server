package message

import (
	"chat-server/internal/model"
	"context"
)

func (s *serv) GetChatMessages(ctx context.Context, chatId int64) ([]model.Message, error) {
	message, err := s.messageRepo.GetChat(ctx, chatId)
	if err != nil {
		return nil, err
	}

	return message, nil
}
