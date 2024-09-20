package converter

import (
	"chat-server/internal/model"
	"chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToMessageFromServ(chat_id int64, message *chat_v1.Message) model.Message {
	return model.Message{
		UserId:   message.GetUserId(),
		ChatId:   chat_id,
		Text:     message.GetText(),
		CreateAt: message.GetTimestamp().AsTime(),
	}
}

func ToMessageFromDesc(message model.Message) *chat_v1.Message {
	return &chat_v1.Message{
		UserId:    message.UserId,
		Text:      message.Text,
		Timestamp: timestamppb.New(message.CreateAt),
	}
}

func ToMessagesFromDesc(messages []model.Message) []*chat_v1.Message {
	result := make([]*chat_v1.Message, len(messages))
	for i := 0; i < len(messages); i++ {
		result[i] = ToMessageFromDesc(messages[i])
	}

	return result
}
