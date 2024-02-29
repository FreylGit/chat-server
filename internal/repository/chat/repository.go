package chat

import (
	"context"
	"github.com/FreylGit/chat-server/internal/client/db"
	"github.com/FreylGit/chat-server/internal/model"
	"github.com/FreylGit/chat-server/internal/repository"
	"github.com/FreylGit/chat-server/internal/repository/chat/converter"
	chtRepo "github.com/FreylGit/chat-server/internal/repository/chat/model"
	sq "github.com/Masterminds/squirrel"
)

const (
	tableName   = "chats"
	idColum     = "id"
	titleColumn = "title"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Chat, error) {
	builder := sq.Select(idColum).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColum: id}).
		Limit(1)
	quety, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	q := db.Query{
		Name:     "chat_repository.Get",
		QueryRaw: quety,
	}
	var model chtRepo.Chat
	err = r.db.DB().ScanOneContext(ctx, &model, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToChatFromRepo(model), err
}

func (r *repo) Create(ctx context.Context, chat *model.Chat) (int64, error) {
	builder := sq.Insert(tableName).
		Columns("id").
		Values(sq.Expr("DEFAULT")).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id")
	query, args, err := builder.ToSql()

	if err != nil {
		return 0, err
	}
	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}
	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).
		Where(sq.Eq{idColum: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRaw: query,
	}
	_, err = r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
