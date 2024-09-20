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

func TestGetChatMessage(t *testing.T) {
	type args struct {
		ctx    context.Context
		chatId int64
	}
	var (
		ctx      = context.Background()
		mc       = minimock.NewController(t)
		chatId   = gofakeit.Int64()
		messages = []model.Message{
			{
				ChatId:   chatId,
				UserId:   gofakeit.Int64(),
				Text:     gofakeit.BeerName(),
				CreateAt: gofakeit.Date(),
			},
			{
				ChatId:   chatId,
				UserId:   gofakeit.Int64(),
				Text:     gofakeit.BeerName(),
				CreateAt: gofakeit.Date(),
			},
			{
				ChatId:   chatId,
				UserId:   gofakeit.Int64(),
				Text:     gofakeit.BeerName(),
				CreateAt: gofakeit.Date(),
			},
		}
	)

	tests := []struct {
		name                  string
		args                  args
		want                  []model.Message
		err                   error
		messageRepositoryMock messageRepositoryMock
	}{
		{
			name: "success case",
			args: args{
				ctx:    ctx,
				chatId: chatId,
			},
			err:  nil,
			want: messages,
			messageRepositoryMock: func(mc *minimock.Controller) repository.MessageRepository {
				mock := repoMock.NewMessageRepositoryMock(mc)
				mock.GetChatMock.Return(messages, nil)
				return mock
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			messageRepo := tt.messageRepositoryMock(mc)
			service := messageServ.NewService(nil, messageRepo)
			result, err := service.GetChatMessages(ctx, tt.args.chatId)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
