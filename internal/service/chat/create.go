package chat

import (
	"context"
	"github.com/FreylGit/chat-server/internal/model"
)

func (s *serv) Create(ctx context.Context, usernames []string) (int64, error) {
	var chatId int64
	chatId, err := s.chatRepository.Create(ctx, &model.Chat{})
	if err != nil {
		return 0, err
	}
	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		if errTx != nil {
			return errTx
		}

		for _, username := range usernames {
			user, err := s.userRepository.GetByName(ctx, username)
			if err != nil {
				userId, err := s.userRepository.Create(ctx, &model.User{Username: username})
				user = &model.User{Id: userId, Username: username}
				if err != nil {
					return err
				}
			}
			err = s.chatUserRepository.Create(ctx, &model.ChatUser{
				UserId: user.Id,
				ChatId: chatId,
			})
			if err != nil {
				return err
			}
		}

		if errTx != nil {
			return errTx
		}

		if errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return chatId, nil
}
