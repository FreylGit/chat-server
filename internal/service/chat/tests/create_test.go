package tests

import (
	moc "chat-server/internal/generate/mocks"
	"chat-server/internal/repository"
	repoMock "chat-server/internal/repository/mocks"
	chatServ "chat-server/internal/service/chat"
	"context"
	"github.com/FreylGit/platform_common/pkg/db"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreate(t *testing.T) {
	type args struct {
		ctx context.Context
		ids []int64
	}
	var (
		ctx    = context.Background()
		mc     = minimock.NewController(t)
		chatId = gofakeit.Int64()
		ids    = RandomId(5)
	)
	//TODO дочистить этот код
	//TODO добавить мок на дургой репозиторий
	txManagerMock := moc.NewTxManagerMock(mc)
	txManagerMock.ReadCommittedMock.Set(func(ctx context.Context, f db.Handler) error {
		return f(ctx)
	})
	tests := []struct {
		name                   string
		args                   args
		err                    error
		want                   int64
		chatRepositoryMock     chatRepositoryMock
		chatUserRepositoryMock chatUserRepositoryMock
		dbMock                 dbMock
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				ids: ids,
			},
			err:  nil,
			want: chatId,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMock.NewChatRepositoryMock(mc)
				mock.CreateMock.Return(chatId, nil)
				return mock
			},
			chatUserRepositoryMock: func(mc *minimock.Controller) repository.ChatUserRepository {
				mock := repoMock.NewChatUserRepositoryMock(mc)
				mock.CreateMock.Return(nil)
				return mock
			},
			dbMock: func(mc *minimock.Controller) db.TxManager {
				mock := moc.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			chatRepo := tt.chatRepositoryMock(mc)
			txManager := tt.dbMock(mc)
			chatUserRepo := tt.chatUserRepositoryMock(mc)
			service := chatServ.NewChatService(txManager, chatRepo, chatUserRepo)
			newChatId, err := service.Create(tt.args.ctx, tt.args.ids)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newChatId)
		})
	}
}
