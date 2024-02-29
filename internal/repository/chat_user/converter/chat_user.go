package converter

import (
	"github.com/FreylGit/chat-server/internal/model"
	modelRepo "github.com/FreylGit/chat-server/internal/repository/chat_user/model"
)

func ToChatUserRepo(chatUser modelRepo.ChatUser) *model.ChatUser {
	return &model.ChatUser{
		ChatId: chatUser.ChatId,
		UserId: chatUser.UserId,
	}
}
