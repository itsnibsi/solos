package main

import (
	"log"

	"github.com/itsnibsi/solos/internal/config"
	"github.com/itsnibsi/solos/internal/logger"
	"github.com/itsnibsi/solos/internal/server"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	logger.Init(cfg.Env)
	defer logger.Sync()

	srv := server.New(cfg)

	logger.Log.Info("Starting server...", zap.String("env", cfg.Env), zap.String("host", cfg.Host), zap.Int("port", cfg.Port))
	if err := srv.Start(); err != nil {
		logger.Log.Error("Error starting server", zap.Error(err))
	}
}
