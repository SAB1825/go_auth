package app

import (
	"context"
	"log"
	"net/http"

	"github.com/SAB1825/go_auth/internal/config"
	"github.com/go-playground/validator"
	"go.uber.org/zap"
)

type AppInterface interface {
	Run()
}

type App struct {
	ServerConfig *config.ServerConfig
	DbConfig     *config.DBConfig
}

func (app *App) Run() {
	ctx := context.Background()

	validate := validator.New()

	if err := validate.Struct(app.ServerConfig); err != nil {
		log.Fatalf("Server config validation failed: %v", err)
	}

	if err := validate.Struct(app.DbConfig); err != nil {
		log.Fatalf("DB config validation failed: %v", err)
	}

	logger := config.GetLogger(app.ServerConfig.AppEnv)

	pool, err := config.SetupDb(ctx, app.DbConfig, logger)

	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}

	defer pool.Close()

	server := &http.Server{
		Addr:         app.ServerConfig.Port,
		Handler:      nil,
		ReadTimeout:  app.ServerConfig.ReadTimeout,
		WriteTimeout: app.ServerConfig.WriteTimeout,
		IdleTimeout:  app.ServerConfig.IdleTimeout,
	}

	logger.Info("Starting the server", zap.String("port", app.ServerConfig.Port))

	err = server.ListenAndServe()
	if err != nil {
		logger.Error("Failed to start the server", zap.Error(err))
	}
}


