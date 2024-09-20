package app

import (
	"chat-server/internal/service"
	"chat-server/internal/service/chat"
	"chat-server/internal/service/message"
	"context"
)

func (sp *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if sp.chatServ == nil {
		sp.chatServ = chat.NewChatService(sp.TxManager(ctx), sp.ChatRepository(ctx), sp.ChatUserRepository(ctx), sp.MessageService(ctx))
	}

	return sp.chatServ
}

func (sp *serviceProvider) MessageService(ctx context.Context) service.MessageService {
	if sp.messageServ == nil {
		sp.messageServ = message.NewService(sp.TxManager(ctx), sp.MessageRepository(ctx))
	}

	return sp.messageServ
}
