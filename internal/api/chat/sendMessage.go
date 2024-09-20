package chat

import (
	"chat-server/internal/converter"
	"chat-server/pkg/chat_v1"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) SendMessage(ctx context.Context, desv *chat_v1.SendMessageRequest) (*emptypb.Empty, error) {
	messServ := converter.ToMessageFromServ(desv.GetChatId(), desv.GetMessage())
	err := i.messageServ.Create(ctx, messServ)
	err = i.chatServ.SendMessage(ctx, messServ)
	return &emptypb.Empty{}, err
}
