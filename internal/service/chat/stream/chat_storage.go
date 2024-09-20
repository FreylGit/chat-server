package stream

import (
	"chat-server/internal/converter"
	"chat-server/internal/model"
	"chat-server/internal/service"
	desc "chat-server/pkg/chat_v1"
	"sync"
)

type storage struct {
	murw  sync.RWMutex
	chats map[int64]*ChatCh
}

func NewChatStorage() service.ChatStorage {
	return &storage{
		chats: make(map[int64]*ChatCh),
	}
}
func (s *storage) ConnectChat(userId int64, chatId int64, stream desc.ChatV1_ConnectChatServer) error {
	_, ok := s.chats[chatId]
	if !ok {
		s.murw.Lock()
		s.chats[chatId] = NewChatCh()
		s.chats[chatId].streams[userId] = stream
		s.murw.Unlock()
		return nil
	}
	s.murw.Lock()
	s.chats[chatId].streams[userId] = stream
	s.murw.Unlock()
	return nil
}

func (s *storage) DisconnectChat(userId int64, chatId int64) error {
	delete(s.chats[chatId].streams, userId)
	return nil
}

func (s *storage) NotifyMessage(chatId int64, message model.Message) {
	chat, ok := s.chats[chatId]
	if !ok {
		s.chats[chatId] = NewChatCh()
	}

	s.chats[chatId].messages <- message
	if len(s.chats[chatId].streams) == 0 {
		return
	}
	s.chats[chatId].murw.RLock()
	for _, user := range chat.streams {
		msg := converter.ToMessageFromDesc(message)
		user.Send(msg)
	}
	s.chats[chatId].murw.RUnlock()
}

type ChatCh struct {
	murw     sync.RWMutex
	messages chan model.Message
	streams  map[int64]desc.ChatV1_ConnectChatServer //мапа стримов пользователей
}

func NewChatCh() *ChatCh {
	return &ChatCh{
		messages: make(chan model.Message, 100),
		streams:  make(map[int64]desc.ChatV1_ConnectChatServer),
	}
}

func (c *ChatCh) SendMessage(message model.Message) {
	c.murw.Lock()
	defer c.murw.Unlock()
	c.messages <- message
}
