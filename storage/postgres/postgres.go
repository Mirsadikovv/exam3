package postgres

import (
	"context"
	"fmt"
	"login/config"
	"login/storage"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/lib/pq"
)

type Store struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, cfg config.Config) (storage.IStorage, error) {
	url := fmt.Sprintf(`host=%s port=%v user=%s password=%s database=%s sslmode=disable`,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

	pgxPoolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}
	pgxPoolConfig.MaxConns = 50
	pgxPoolConfig.MaxConnLifetime = time.Hour

	newPool, err := pgxpool.NewWithConfig(ctx, pgxPoolConfig)
	if err != nil {
		return nil, err
	}

	return Store{
		Pool: newPool,
	}, nil
}

func (s Store) CloseDB() {
	s.Pool.Close()
}

func (s Store) CustomerStorage() storage.CustomerStorage {
	newCustomer := NewCustomer(s.Pool)

	return &newCustomer
}
