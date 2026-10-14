package main

import (
	"log"

	"github.com/SAB1825/go_auth/cmd/app"
	"github.com/SAB1825/go_auth/internal/config"
)

func main() {
	serverConfig, err := config.LoadServerConfig()
	if err != nil {
		log.Fatal("Failed to load the server config: " + err.Error())
	}
	dbConfig, err := config.LoadDbConfig()
	if err != nil {
		log.Fatal("Failed to load the DB config: " + err.Error())
	}

	application := &app.App{
		ServerConfig: serverConfig,
		DbConfig:     dbConfig,
	}
	application.Run()
}
