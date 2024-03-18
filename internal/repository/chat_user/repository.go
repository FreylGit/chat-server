package chat_user

import (
	"context"
	"github.com/FreylGit/chat-server/internal/model"
	"github.com/FreylGit/chat-server/internal/repository"
	"github.com/FreylGit/platform_common/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

const (
	tableName    = "chat_users"
	chatIdColumn = "chat_id"
	userIdColumn = "user_id"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ChatUserRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, chatUser *model.ChatUser) error {
	builder := sq.Insert(tableName).
		Columns(chatIdColumn, userIdColumn).
		Values(chatUser.ChatId, chatUser.UserId).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "chatUser_repository.Create",
		QueryRaw: query,
	}
	_, err = r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Contains(ctx context.Context, chatUser *model.ChatUser) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repo) Delete(ctx context.Context, chatUser *model.ChatUser) error {
	//TODO implement me
	panic("implement me")
}
