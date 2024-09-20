package repository

import (
	"chat-server/internal/model"
	"context"
)

type ChatRepository interface {
	Create(ctx context.Context) (int64, error)
	Delete(ctx context.Context, id int64) error
}

type ChatUserRepository interface {
	Create(ctx context.Context, chat_id int64, ids []int64) error
	Delete(ctx context.Context, chat_id int64, user_id int64) error
}

type MessageRepository interface {
	Create(ctx context.Context, model model.Message) error
	Delete(ctx context.Context, id int64, user_id int64) error
	GetChat(ctx context.Context, chat_id int64) ([]model.Message, error)
}
