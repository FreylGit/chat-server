package chat

import (
	desc "chat-server/pkg/chat_v1"
)

func (i *Implementation) ConnectChat(req *desc.ConnectChatRequest, stream desc.ChatV1_ConnectChatServer) error {
	err := i.chatServ.ConnectChat(req.GetChatId(), req.GetUserId(), stream)
	return err
}
