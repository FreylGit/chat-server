package converter

import (
	"github.com/FreylGit/chat-server/internal/model"
	modelRepo "github.com/FreylGit/chat-server/internal/repository/message/model"
)

func ToMessageRepo(message modelRepo.Message) *model.Message {
	return &model.Message{
		Id:     message.Id,
		ChatId: message.ChatId,
		UserId: message.UserId,
	}
}
