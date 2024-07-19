package main

import (
	"log"

	"github.com/itsnibsi/solos/internal/config"
	"github.com/itsnibsi/solos/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	srv := server.New(cfg)

	if err := srv.Start(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
