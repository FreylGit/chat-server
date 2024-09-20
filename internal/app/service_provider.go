package app

import (
	"chat-server/internal/api/chat"
	"chat-server/internal/config/env"
	"chat-server/internal/logger"
	"chat-server/internal/repository"
	"chat-server/internal/service"
	"context"
	"github.com/FreylGit/platform_common/pkg/closer"
	"github.com/FreylGit/platform_common/pkg/db"
	"github.com/FreylGit/platform_common/pkg/db/pg"
	"github.com/FreylGit/platform_common/pkg/db/transaction"
	"go.uber.org/zap"
)

type serviceProvider struct {
	grpcConfig         env.GRPCConfig
	pgConfig           env.PGConfig
	chatImpl           *chat.Implementation
	dbClient           db.Client
	txManager          db.TxManager
	chatServ           service.ChatService
	messageServ        service.MessageService
	chatRepository     repository.ChatRepository
	messageRepository  repository.MessageRepository
	chatUserRepository repository.ChatUserRepository
}

func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) ChatImpl(ctx context.Context) *chat.Implementation {
	if sp.chatImpl == nil {
		sp.chatImpl = chat.NewImplementation(sp.ChatService(ctx), sp.MessageService(ctx))
	}

	return sp.chatImpl
}

func (sp *serviceProvider) GrpcConfig() env.GRPCConfig {
	if sp.grpcConfig == nil {
		config, err := env.NewGRPCConfig()
		if err != nil {
			logger.Fatal("grpc config err:", zap.Error(err))
		}
		sp.grpcConfig = config
	}

	return sp.grpcConfig
}

func (sp *serviceProvider) PgConfig() env.PGConfig {
	if sp.pgConfig == nil {
		config, err := env.NewPGConfig()
		if err != nil {
			logger.Fatal("grpc config err: ", zap.Error(err))
		}
		sp.pgConfig = config
	}

	return sp.pgConfig
}

func (sp *serviceProvider) DbClient(ctx context.Context) db.Client {
	if sp.dbClient == nil {
		cl, err := pg.New(ctx, sp.PgConfig().DSN())
		if err != nil {
			logger.Fatal("grpc config err:", zap.Error(err))
		}
		err = cl.DB().Ping(ctx)
		if err != nil {
			logger.Fatal("grpc config err:", zap.Error(err))
		}
		closer.Add(cl.Close)
		sp.dbClient = cl
	}

	return sp.dbClient
}

func (sp *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if sp.txManager == nil {
		sp.txManager = transaction.NewTransactionManager(sp.DbClient(ctx).DB())
	}

	return sp.txManager
}
