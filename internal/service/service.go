package service

import (
	"context"
	"github.com/FreylGit/chat-server/internal/model"
)

type ChatService interface {
	Create(ctx context.Context, usernames []string) (int64, error)
	Delete(ctx context.Context, id int64) error
	SendMessage(ctx context.Context, message model.Message) error
}
