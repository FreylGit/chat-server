package chat

import (
	desc "chat-server/pkg/chat_v1"
	"sync"
)

func (s *serv) ConnectChat(chatId int64, userId int64, stream desc.ChatV1_ConnectChatServer) error {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.chatStorage.ConnectChat(userId, chatId, stream)
	}()
	wg.Wait()

	for {
		select {
		case <-stream.Context().Done():
			go s.chatStorage.DisconnectChat(userId, chatId)
			return nil
		}
	}
}
