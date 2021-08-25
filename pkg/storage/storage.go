package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kbats183/CTStickersBot/pkg/core"
)

type Storage struct {
	config     core.StorageConfig
	clientPull *pgxpool.Pool
}

func NewStorage(ctx context.Context, config core.StorageConfig) (*Storage, error) {
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
