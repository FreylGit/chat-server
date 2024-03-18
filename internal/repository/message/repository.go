package message

import (
	"context"
	"github.com/FreylGit/chat-server/internal/model"
	"github.com/FreylGit/chat-server/internal/repository"
	"github.com/FreylGit/platform_common/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

const (
	tableName       = "messages"
	idColumn        = "id"
	chatIdColumn    = "chat_id"
	userIdColumn    = "user_id"
	textColumn      = "text"
	timestampColumn = "timestamp"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.MessageRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Message, error) {
	builder := sq.Select(idColumn, chatIdColumn, userIdColumn, textColumn, timestampColumn).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "message_repository.Get",
		QueryRaw: query,
	}
	var message model.Message
	err = r.db.DB().ScanOneContext(ctx, &message, q, args...)

	return &message, err
}

func (r *repo) Create(ctx context.Context, message *model.Message) (int64, error) {
	builder := sq.Insert(tableName).
		Columns(chatIdColumn, userIdColumn, textColumn, timestampColumn).
		Values(message.ChatId, message.UserId, message.Text, message.Timestamp).
		PlaceholderFormat(sq.Dollar)
	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}
	q := db.Query{
		Name:     "message_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
