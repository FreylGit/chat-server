package user

import (
	"context"
	"github.com/FreylGit/chat-server/internal/model"
	"github.com/FreylGit/chat-server/internal/repository"
	"github.com/FreylGit/chat-server/internal/repository/user/converter"
	userRepo "github.com/FreylGit/chat-server/internal/repository/user/model"
	"github.com/FreylGit/platform_common/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

const (
	tableName      = "users"
	idColumn       = "id"
	usernameColumn = "username"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.Select(idColumn, usernameColumn).
		From(tableName).Where(sq.Eq{idColumn: id}).
		PlaceholderFormat(sq.Dollar)
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}
	var user userRepo.User
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)

	return converter.ToUserFromRepo(user), nil
}

func (r *repo) GetByName(ctx context.Context, username string) (*model.User, error) {
	builder := sq.Select(idColumn, usernameColumn).
		From(tableName).Where(sq.Eq{usernameColumn: username}).
		PlaceholderFormat(sq.Dollar)
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	q := db.Query{
		Name:     "user_repository.GetByName",
		QueryRaw: query,
	}
	var user userRepo.User
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(user), nil
}

func (r *repo) Create(ctx context.Context, user *model.User) (int64, error) {
	builder := sq.Insert(tableName).
		Columns(usernameColumn).
		Values(user.Username).
		PlaceholderFormat(sq.Dollar).Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}
	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}
	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repo) Exist(ctx context.Context, userId int64) (bool, error) {
	builder := sq.Select(idColumn, usernameColumn).
		From(tableName).Where(sq.Eq{idColumn: userId}).
		PlaceholderFormat(sq.Dollar)
	query, args, err := builder.ToSql()
	if err != nil {
		return false, err
	}
	q := db.Query{
		Name:     "user_repository.Exist",
		QueryRaw: query,
	}
	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return false, err
	}
	if id == 0 {
		return false, nil
	}

	return true, nil
}
