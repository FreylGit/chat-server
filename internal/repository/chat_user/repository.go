package chat_user

import (
	"chat-server/internal/repository"
	"context"
	"fmt"
	"github.com/FreylGit/platform_common/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

const (
	tableName    = "chat_user"
	userIdColumn = "user_id"
	chatIdColumn = "chat_id"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ChatUserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, chat_id int64, ids []int64) error {
	builder := sq.Insert(tableName).
		Columns(userIdColumn, chatIdColumn).PlaceholderFormat(sq.Dollar)
	for _, id := range ids {
		builder = builder.Values(id, chat_id)
	}
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "chat_user_repository.Create",
		QueryRow: query,
	}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, chat_id int64, user_id int64) error {
	builder := sq.Delete(tableName).
		Where(sq.And{sq.Eq{chatIdColumn: chat_id}, sq.Eq{userIdColumn: user_id}}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "chat_user_repository.Delete",
		QueryRow: query,
	}

	result, err := r.db.DB().ExecContext(ctx, q, args...)
	if result.RowsAffected() == 0 {
		return fmt.Errorf("chat_user_repository.Delete: chat_id=%d user_id=%d", chat_id, user_id)
	}

	return err
}
