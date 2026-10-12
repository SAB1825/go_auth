package config

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port         string        `validate:"required"`
	ReadTimeout  time.Duration `validate:"required"`
	WriteTimeout time.Duration `validate:"required"`
	IdleTimeout  time.Duration `validate:"required"`
	AppEnv       string        `validate:"required"`
}

func LoadServerConfig() (*ServerConfig, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Something went wrong when loading the env file: " + err.Error())
		return nil, err
	}

	cfg := &ServerConfig{
		Port:         LoadSingleEnvVar("PORT", ":3001"),
		ReadTimeout:  time.Duration(LoadSingleEnvVar("READ_TIMEOUT", 15)) * time.Second,
		WriteTimeout: time.Duration(LoadSingleEnvVar("WRITE_TIMEOUT", 15)) * time.Second,
		IdleTimeout:  time.Duration(LoadSingleEnvVar("IDLE_TIMEOUT", 150)) * time.Second,
		AppEnv:       LoadSingleEnvVar("APP_ENV", "development"),
	}

	return cfg, nil
}
