package message

import (
	modelServ "chat-server/internal/model"
	"chat-server/internal/repository"
	"chat-server/internal/repository/message/converter"
	"chat-server/internal/repository/message/model"
	"context"
	"fmt"
	"github.com/FreylGit/platform_common/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

const (
	tableName      = "message"
	idColumn       = "id"
	userIdColumn   = "user_id"
	chatIdColumn   = "chat_id"
	textColumn     = "text"
	createAtColumn = "create_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.MessageRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, model modelServ.Message) error {
	builder := sq.Insert(tableName).
		Columns(userIdColumn, chatIdColumn, textColumn).
		Values(model.UserId, model.ChatId, model.Text).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "message_repository.Create",
		QueryRow: query,
	}

	result, err := r.db.DB().ExecContext(ctx, q, args...)

	if result.RowsAffected() != 1 {
		return fmt.Errorf("insert message failed, rows affected: %d", result.RowsAffected())
	}

	return err
}

func (r *repo) Delete(ctx context.Context, id int64, user_id int64) error {
	builder := sq.Delete(tableName).
		Where(sq.And{sq.Eq{idColumn: id}, sq.Eq{userIdColumn: user_id}}).
		PlaceholderFormat(sq.Dollar)
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "message_repository.Delete",
		QueryRow: query,
	}
	result, err := r.db.DB().ExecContext(ctx, q, args...)
	if result.RowsAffected() != 1 {
		return fmt.Errorf("delete message failed, rows affected: %d", result.RowsAffected())
	}

	return err
}

func (r *repo) GetChat(ctx context.Context, chat_id int64) ([]modelServ.Message, error) {
	builder := sq.Select(idColumn, userIdColumn, chatIdColumn, textColumn, createAtColumn).
		Where(sq.Eq{chatIdColumn: chat_id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	q := db.Query{
		Name:     "message_repository.GetChat",
		QueryRow: query,
	}
	var messages []model.Message
	err = r.db.DB().ScanAllContext(ctx, &messages, q, args...)

	return converter.ToMessagesFromRepo(messages), err
}
