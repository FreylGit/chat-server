package chat

import (
	"context"
	"github.com/FreylGit/platform_common/pkg/db"

	"github.com/FreylGit/chat-server/internal/model"
	"github.com/FreylGit/chat-server/internal/repository"
	"github.com/FreylGit/chat-server/internal/service"
)

type serv struct {
	chatRepository        repository.ChatRepository
	chatMessageRepository repository.MessageRepository
	chatUserRepository    repository.ChatUserRepository
	userRepository        repository.UserRepository
	txManager             db.TxManager
}

func NewService(chatRepository repository.ChatRepository,
	chatMessageRepository repository.MessageRepository,
	chatUserRepository repository.ChatUserRepository,
	userRepository repository.UserRepository,
	txManager db.TxManager) service.ChatService {

	return &serv{
		chatRepository:        chatRepository,
		chatMessageRepository: chatMessageRepository,
		chatUserRepository:    chatUserRepository,
		userRepository:        userRepository,
		txManager:             txManager,
	}
}

func (s *serv) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (s *serv) SendMessage(ctx context.Context, message model.Message) error {
	//TODO implement me
	panic("implement me")
}
