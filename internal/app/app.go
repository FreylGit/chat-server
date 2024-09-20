package app

import (
	"chat-server/internal/interceptor"
	"chat-server/internal/logger"
	"chat-server/internal/metric"
	"chat-server/internal/rate_limiter"
	"chat-server/internal/tracing"
	"chat-server/pkg/chat_v1"
	"context"
	"github.com/natefinch/lumberjack"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

type App struct {
	serviceProvider  *serviceProvider
	grpcServer       *grpc.Server
	prometheusServer *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := a.runGRPCServer(ctx)
		if err != nil {
			logger.Fatal("grpc server error: %v", zap.Error(err))
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := a.runPrometheus()
		if err != nil {
			logger.Fatal("Prometheus server error:", zap.Error(err))
		}
	}()
	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx2 context.Context) error{
		a.initServiceProvider,
		a.initGRPCServer,
		metric.Init,
	}
	for _, init := range inits {
		err := init(ctx)
		if err != nil {
			return err
		}
	}
	logger.Init(getCore(getAtomicLevel()))
	tracing.Init(logger.Logger(), "chat-service")
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	rateLimiter := rate_limiter.NewTokenBucketLimiter(ctx, 10, time.Second)
	a.grpcServer = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.LogInterceptor,
			interceptor.ServerTracingInterceptor,
			interceptor.NewRateLimiterInterceptor(rateLimiter).Unary,
			interceptor.MetricsInterceptor,
		))
	reflection.Register(a.grpcServer)

	chat_v1.RegisterChatV1Server(
		a.grpcServer,
		a.serviceProvider.ChatImpl(ctx),
	)

	return nil
}

func (a *App) initServiceProvider(ctx context.Context) error {
	a.serviceProvider = NewServiceProvider()

	return nil
}

func (a *App) runGRPCServer(ctx context.Context) error {
	lis, err := net.Listen("tcp", a.serviceProvider.GrpcConfig().Address())
	if err != nil {
		return err
	}
	err = a.grpcServer.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runPrometheus() error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	server := &http.Server{
		Addr:    "localhost:2112",
		Handler: mux,
	}
	a.prometheusServer = server
	log.Printf("Prometheus server is running on localhost:2112")
	err := a.prometheusServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func getCore(level zap.AtomicLevel) zapcore.Core {
	stdout := zapcore.AddSync(os.Stdout)
	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10, // megabyte
		MaxBackups: 3,
		MaxAge:     7, //days
	})
	prodactionCfg := zap.NewProductionEncoderConfig()
	prodactionCfg.TimeKey = "timestamp"
	prodactionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
	fileEncoder := zapcore.NewJSONEncoder(prodactionCfg)

	return zapcore.NewTee(zapcore.NewCore(consoleEncoder, stdout, level),
		zapcore.NewCore(fileEncoder, file, level))
}
func getAtomicLevel() zap.AtomicLevel {
	var level zapcore.Level
	err := level.Set("debug")
	if err != nil {
		log.Fatalf("failed to set log level: %v", err)
	}
	return zap.NewAtomicLevelAt(level)
}
