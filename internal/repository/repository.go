package repository

import (
	"context"
	"github.com/FreylGit/chat-server/internal/model"
)

type ChatRepository interface {
	Get(ctx context.Context, id int64) (*model.Chat, error)
	Create(ctx context.Context, chat *model.Chat) (int64, error)
	Delete(ctx context.Context, id int64) error
}

type MessageRepository interface {
	Get(ctx context.Context, id int64) (*model.Message, error)
	Create(ctx context.Context, message *model.Message) (int64, error)
}

type UserRepository interface {
	Get(ctx context.Context, id int64) (*model.User, error)
	GetByName(ctx context.Context, username string) (*model.User, error)
	Create(ctx context.Context, user *model.User) (int64, error)
	Exist(ctx context.Context, userId int64) (bool, error)
}

type ChatUserRepository interface {
	Create(ctx context.Context, chatUser *model.ChatUser) error
	Contains(ctx context.Context, chatUser *model.ChatUser) (bool, error)
	Delete(ctx context.Context, chatUser *model.ChatUser) error
}
