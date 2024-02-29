package app

import (
	"context"
	chatImp "github.com/FreylGit/chat-server/internal/api/chat"
	"github.com/FreylGit/chat-server/internal/client/db"
	"github.com/FreylGit/chat-server/internal/client/db/pg"
	"github.com/FreylGit/chat-server/internal/client/db/transaction"
	"github.com/FreylGit/chat-server/internal/closer"
	"github.com/FreylGit/chat-server/internal/config"
	"github.com/FreylGit/chat-server/internal/repository"
	"github.com/FreylGit/chat-server/internal/repository/chat"
	"github.com/FreylGit/chat-server/internal/repository/chat_user"
	"github.com/FreylGit/chat-server/internal/repository/message"
	"github.com/FreylGit/chat-server/internal/repository/user"
	"github.com/FreylGit/chat-server/internal/service"
	chatService "github.com/FreylGit/chat-server/internal/service/chat"
	"log"
)

type serviceProvider struct {
	dbClient           db.Client
	grpcConfig         config.GRPCConfig
	pgConfig           config.PGConfig
	chatRepository     repository.ChatRepository
	chatUserRepository repository.ChatUserRepository
	messageRepository  repository.MessageRepository
	userRepository     repository.UserRepository
	chatImpl           *chatImp.Implementation
	chatService        service.ChatService
	txManager          db.TxManager
}

func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DNS())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chat.NewRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) ChatMessageRepository(ctx context.Context) repository.MessageRepository {
	if s.messageRepository == nil {
		s.messageRepository = message.NewRepository(s.DBClient(ctx))
	}

	return s.messageRepository
}

func (s *serviceProvider) CharUserRepository(ctx context.Context) repository.ChatUserRepository {
	if s.chatUserRepository == nil {
		s.chatUserRepository = chat_user.NewRepository(s.DBClient(ctx))
	}

	return s.chatUserRepository
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = user.NewRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) ChatImplementation(ctx context.Context) *chatImp.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chatImp.NewImplementation(s.ChatService(ctx))
	}

	return s.chatImpl
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(s.ChatRepository(ctx), s.ChatMessageRepository(ctx), s.ChatUserRepository(ctx), s.UserRepository(ctx), s.TxManager(ctx))
	}

	return s.chatService
}

func (s *serviceProvider) ChatUserRepository(ctx context.Context) repository.ChatUserRepository {
	if s.chatUserRepository == nil {
		s.chatUserRepository = chat_user.NewRepository(s.DBClient(ctx))
	}

	return s.chatUserRepository
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed parse grpc config")
		}
		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed parse pg config")
		}
		s.pgConfig = cfg
	}

	return s.pgConfig
}
