package message

import (
	"chat-server/internal/repository"
	"chat-server/internal/service"
	"github.com/FreylGit/platform_common/pkg/db"
)

type serv struct {
	messageRepo repository.MessageRepository
	txManager   db.TxManager
}

func NewService(txManager db.TxManager, messageRepo repository.MessageRepository) service.MessageService {
	return &serv{
		txManager:   txManager,
		messageRepo: messageRepo,
	}
}
