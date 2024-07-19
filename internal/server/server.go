package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/itsnibsi/solos/internal/api/handlers"
	"github.com/itsnibsi/solos/internal/config"
	"github.com/itsnibsi/solos/internal/errors"
)

type Server struct {
	config *config.Config
	app    *fiber.App
}

func New(cfg *config.Config) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: errors.ErrorHandler,
	})

	return &Server{
		config: cfg,
		app:    app,
	}
}

func (s *Server) SetupRoutes() {
	s.app.Get("/health", handlers.Health)
}

func (s *Server) Start() error {
	s.SetupRoutes()

	addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	log.Printf("Starting server on %s", addr)
	return s.app.Listen(addr)
}
