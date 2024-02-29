package converter

import (
	"github.com/FreylGit/chat-server/internal/model"
	modelRepo "github.com/FreylGit/chat-server/internal/repository/chat/model"
)

func ToChatFromRepo(chat modelRepo.Chat) *model.Chat {
	return &model.Chat{
		Id: chat.Id,
	}
}
