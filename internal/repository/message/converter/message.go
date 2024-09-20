package converter

import (
	modelServ "chat-server/internal/model"
	"chat-server/internal/repository/message/model"
)

func ToMessageFromRepo(message model.Message) modelServ.Message {
	return modelServ.Message{
		Id:       message.Id,
		UserId:   message.UserId,
		ChatId:   message.ChatId,
		Text:     message.Text,
		CreateAt: message.CreateAt,
	}
}

func ToMessagesFromRepo(messages []model.Message) []modelServ.Message {
	result := make([]modelServ.Message, len(messages))
	for i, message := range messages {
		result[i] = ToMessageFromRepo(message)
	}

	return result
}

func ToMessageFromServ(message modelServ.Message) model.Message {
	return model.Message{
		Id:       message.Id,
		UserId:   message.UserId,
		ChatId:   message.ChatId,
		Text:     message.Text,
		CreateAt: message.CreateAt,
	}
}
