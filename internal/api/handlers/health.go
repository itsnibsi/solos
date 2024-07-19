package handlers

import (
	"github.com/gofiber/fiber/v3"
)

// Health godoc
// @Summary Health
// @Description Get health status
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func Health(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
