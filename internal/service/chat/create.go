package chat

import (
	"context"
	"fmt"
)

func (s *serv) Create(ctx context.Context, ids []int64) (int64, error) {
	var chatId int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var txErr error
		chatId, txErr = s.chatRepo.Create(ctx)
		if txErr != nil {
			return fmt.Errorf("Ошибка при создании чата: %w", txErr)
		}
		txErr = s.chatUserRepo.Create(ctx, chatId, ids)

		return txErr
	})
	if err != nil {
		return 0, err
	}

	return chatId, nil
}
