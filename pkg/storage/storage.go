package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type StorageConfig struct {
	Host     string `yaml:"host" env:"DB_HOST"`
	UserName string `yaml:"user" env:"DB_USER"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	DBName   string `yaml:"name" env:"DB_NAME"`
}

type Storage struct {
	config     StorageConfig
	clientPull *pgxpool.Pool
}

func NewStorage(ctx context.Context, config StorageConfig) (*Storage, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s", config.UserName, config.Password, config.Host, config.DBName)
	pollConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.ConnectConfig(ctx, pollConfig)
	if err != nil {
		return nil, err
	}
	return &Storage{
		config:     config,
		clientPull: pool,
	}, nil
}
