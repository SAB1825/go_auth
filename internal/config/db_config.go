package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type DBConfig struct {
	DbUrl string `validate:"required"`
}

func LoadDbConfig() (*DBConfig, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Something went wrong when loading the env file: " + err.Error())
		return nil, err
	}

	cfg := &DBConfig{
		DbUrl: LoadSingleEnvVar("DATABASE_URL", "postgresql://postgres:postgres@localhost:5432/go_auth"),
	}

	return cfg, nil
}

func SetupDb(ctx context.Context, cfg *DBConfig, logger *zap.Logger) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, cfg.DbUrl)
	if err != nil {
		logger.Error("unable to create db pool", zap.Error(err))
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		logger.Error("unable to reach database", zap.Error(err))
		return nil, err
	}

	logger.Info("Successfully connected to the database")

	return pool, nil
}
