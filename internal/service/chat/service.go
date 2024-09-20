package chat

import (
	"chat-server/internal/repository"
	"chat-server/internal/service"
	"chat-server/internal/service/chat/stream"
	"github.com/FreylGit/platform_common/pkg/db"
)

type serv struct {
	txManager    db.TxManager
	chatRepo     repository.ChatRepository
	chatUserRepo repository.ChatUserRepository
	chatStorage  service.ChatStorage
	messageServ  service.MessageService
}

func NewChatService(txManager db.TxManager,
	chatRepo repository.ChatRepository,
	chatUserRepo repository.ChatUserRepository,

	messageServ service.MessageService) service.ChatService {
	return &serv{
		txManager:    txManager,
		chatRepo:     chatRepo,
		chatUserRepo: chatUserRepo,
		chatStorage:  stream.NewChatStorage(),
		messageServ:  messageServ,
	}
}
