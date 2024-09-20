package chat

import (
	"chat-server/internal/service"
	"chat-server/pkg/chat_v1"
)

type Implementation struct {
	chat_v1.UnsafeChatV1Server
	chatServ    service.ChatService
	messageServ service.MessageService
}

func NewImplementation(chatServ service.ChatService, messageServ service.MessageService) *Implementation {
	return &Implementation{
		chatServ:    chatServ,
		messageServ: messageServ,
	}
}
