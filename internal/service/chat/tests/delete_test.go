package tests

import (
	"chat-server/internal/repository"
	"chat-server/internal/repository/mocks"
	"chat-server/internal/service/chat"
	"context"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDelete(t *testing.T) {
	type args struct {
		ctx    context.Context
		chatId int64
	}

	var (
		ctx    = context.Background()
		mc     = minimock.NewController(t)
		chatId = gofakeit.Int64()
	)

	test := []struct {
		name               string
		args               args
		err                error
		chatRepositoryMock chatRepositoryMock
	}{
		{
			name: "success case",
			args: args{
				ctx:    ctx,
				chatId: chatId,
			},
			err: nil,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := mocks.NewChatRepositoryMock(mc)
				mock.DeleteMock.Return(nil)
				return mock
			},
		},
	}
	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			chatRepo := tt.chatRepositoryMock(mc)
			service := chat.NewChatService(nil, chatRepo, nil)
			err := service.Delete(tt.args.ctx, tt.args.chatId)

			require.Equal(t, tt.err, err)
		})
	}
}
