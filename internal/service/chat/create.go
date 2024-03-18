package chat

import (
	"context"
	"github.com/FreylGit/chat-server/internal/model"
)

func (s *serv) Create(ctx context.Context, usernames []string) (int64, error) {
	var chatId int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		chatId, errTx = s.chatRepository.Create(ctx, &model.Chat{})
		if errTx != nil {
			return errTx
		}
		for _, username := range usernames {
			user, errTx := s.userRepository.GetByName(ctx, username)
			if errTx != nil {
				var userId int64
				userId, errTx = s.userRepository.Create(ctx, &model.User{Username: username})
				user = &model.User{Id: userId, Username: username}
				if errTx != nil {
					return errTx
				}
			}
			errTx = s.chatUserRepository.Create(ctx, &model.ChatUser{
				UserId: user.Id,
				ChatId: chatId,
			})
			if errTx != nil {
				return errTx
			}
		}

		return nil
	})
	if err != nil {
		return 0, err
	}
	return chatId, nil
}
