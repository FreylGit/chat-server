package app

import (
	"chat-server/internal/repository"
	chat2 "chat-server/internal/repository/chat"
	"chat-server/internal/repository/chat_user"
	"chat-server/internal/repository/message"
	"context"
)

func (sp *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if sp.chatRepository == nil {
		sp.chatRepository = chat2.NewRepository(sp.DbClient(ctx))
	}

	return sp.chatRepository
}

func (sp *serviceProvider) ChatUserRepository(ctx context.Context) repository.ChatUserRepository {
	if sp.chatUserRepository == nil {
		sp.chatUserRepository = chat_user.NewRepository(sp.DbClient(ctx))
	}

	return sp.chatUserRepository
}

func (sp *serviceProvider) MessageRepository(ctx context.Context) repository.MessageRepository {
	if sp.messageRepository == nil {
		sp.messageRepository = message.NewRepository(sp.DbClient(ctx))
	}

	return sp.messageRepository
}
