package main

import (
	"log"

	"github.com/itsnibsi/solos/internal/config"
	"github.com/itsnibsi/solos/internal/server"
)

func main() {
	cfg := config.Load()
	srv := server.New(cfg)

	if err := srv.Start(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
