package converter

import (
	"github.com/FreylGit/chat-server/internal/model"
	desc "github.com/FreylGit/chat-server/pkg/chat_v1"
)

func ToChatFromService(chat *model.Chat) *desc.CreateResponse {
	return &desc.CreateResponse{
		Id: chat.Id,
	}
}

func ToChatFromDesc(chat *desc.CreateRequest) *model.Chat {
	return &model.Chat{}
}
