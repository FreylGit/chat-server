package tests

import (
	"chat-server/internal/repository"
	"github.com/FreylGit/platform_common/pkg/db"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
)

type chatRepositoryMock func(mc *minimock.Controller) repository.ChatRepository
type chatUserRepositoryMock func(mc *minimock.Controller) repository.ChatUserRepository
type dbMock func(mc *minimock.Controller) db.TxManager

func RandomId(n int) []int64 {
	ids := make([]int64, n, n)
	for i := 0; i < n; i++ {
		ids[i] = gofakeit.Int64()
	}

	return ids
}
