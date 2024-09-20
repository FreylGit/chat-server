package env

import (
	"fmt"
	"os"
)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

type GRPCConfig interface {
	Address() string
}

func NewGRPCConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if host == "" {
		return nil, fmt.Errorf("environment variable %s not set", grpcHostEnvName)
	}
	port := os.Getenv(grpcPortEnvName)
	if port == "" {
		return nil, fmt.Errorf("environment variable %s not set", grpcPortEnvName)
	}

	return &grpcConfig{host: host, port: port}, nil
}

type grpcConfig struct {
	host string
	port string
}

func (g grpcConfig) Address() string {
	return fmt.Sprintf("%s:%s", g.host, g.port)
}
