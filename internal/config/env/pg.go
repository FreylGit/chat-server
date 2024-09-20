package env

import (
	"fmt"
	"os"
)

const (
	pgDsnEnvName = "PG_DSN"
)

type PGConfig interface {
	DSN() string
}

func NewPGConfig() (PGConfig, error) {
	dsn := os.Getenv(pgDsnEnvName)
	if dsn == "" {
		return nil, fmt.Errorf("environment variable %s not set", pgDsnEnvName)
	}
	return &pgConfig{dsn: dsn}, nil
}

type pgConfig struct {
	dsn string
}

func (p pgConfig) DSN() string {
	return p.dsn
}
