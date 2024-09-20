package tests

import (
	"chat-server/internal/repository"
	"github.com/gojuno/minimock/v3"
)

type messageRepositoryMock func(mc *minimock.Controller) repository.MessageRepository
