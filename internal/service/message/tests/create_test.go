package tests

import (
	"chat-server/internal/model"
	"chat-server/internal/repository"
	repoMock "chat-server/internal/repository/mocks"
	messageServ "chat-server/internal/service/message"
	"context"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreate(t *testing.T) {
	type args struct {
		ctx     context.Context
		message model.Message
	}

	var (
		ctx     = context.Background()
		mc      = minimock.NewController(t)
		message = model.Message{
			UserId: gofakeit.Int64(),
			ChatId: gofakeit.Int64(),
			Text:   gofakeit.BeerName(),
		}
	)
	tests := []struct {
		name                  string
		args                  args
		err                   error
		messageRepositoryMock messageRepositoryMock
	}{
		{
			name: "success case",
			args: args{
				ctx:     ctx,
				message: message,
			},
			err: nil,
			messageRepositoryMock: func(mc *minimock.Controller) repository.MessageRepository {
				mock := repoMock.NewMessageRepositoryMock(mc)
				mock.CreateMock.Return(nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			messageRepo := tt.messageRepositoryMock(mc)
			service := messageServ.NewService(nil, messageRepo)
			err := service.Create(ctx, tt.args.message)
			require.Equal(t, tt.err, err)
		})
	}
}
