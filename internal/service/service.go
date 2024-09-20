package service

import (
	"chat-server/internal/model"
	desc "chat-server/pkg/chat_v1"
	"context"
)

type ChatService interface {
	Create(ctx context.Context, ids []int64) (int64, error)
	Delete(ctx context.Context, id int64) error
	ConnectChat(chatId int64, userId int64, stream desc.ChatV1_ConnectChatServer) error
	SendMessage(ctx context.Context, message model.Message) error
}

type MessageService interface {
	Create(ctx context.Context, message model.Message) error
	GetChatMessages(ctx context.Context, chatId int64) ([]model.Message, error)
}

type ChatStorage interface {
	ConnectChat(userId int64, chatId int64, stream desc.ChatV1_ConnectChatServer) error
	DisconnectChat(userId int64, chatId int64) error
	NotifyMessage(chatId int64, message model.Message)
}
