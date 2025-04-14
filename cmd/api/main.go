package main

import (
	"fmt"

	"banking-service/internal/api"
	"banking-service/internal/config"
	"banking-service/internal/db"
	"banking-service/internal/repository"
	"banking-service/internal/service"
	"banking-service/pkg/logger"
)

func main() {
	cfg := config.Load()

	database, err := db.Connect(cfg)
	if err != nil {
		logger.Critical(fmt.Sprintf("Failed to connect to database: %v", err))
		return
	}

	nasabahRepo := repository.NewNasabahRepository(database)
	accountService := service.NewAccountService(nasabahRepo)

	router := api.SetupRouter(accountService)

	serverAddr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	logger.Info(fmt.Sprintf("Starting server on %s", serverAddr))

	if err := router.Start(serverAddr); err != nil {
		logger.Critical(fmt.Sprintf("Failed to start server: %v", err))
	}
}
