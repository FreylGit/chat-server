package converter

import (
	"github.com/FreylGit/chat-server/internal/model"
	modelRepo "github.com/FreylGit/chat-server/internal/repository/user/model"
)

func ToUserFromRepo(user modelRepo.User) *model.User {
	return &model.User{
		Id:       user.Id,
		Username: user.Username,
	}
}
