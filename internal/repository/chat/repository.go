package chat

import (
	"chat-server/internal/repository"
	"context"
	"fmt"
	"github.com/FreylGit/platform_common/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

const (
	tableName = "chat"
	idColumn  = "id"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context) (int64, error) {
	querySql := "INSERT INTO chat DEFAULT VALUES RETURNING id"

	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRow: querySql,
	}
	var id int64
	err := r.db.DB().QueryRowContext(ctx, q).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).
		Where(sq.Eq{idColumn: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRow: query,
	}

	result, err := r.db.DB().ExecContext(ctx, q, args...)

	if result.RowsAffected() == 0 {
		return fmt.Errorf("chat_repository.Delete: no such id: %d", id)
	}

	return err
}
