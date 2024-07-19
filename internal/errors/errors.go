package errors

import "github.com/gofiber/fiber/v3"

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e AppError) Error() string {
	return e.Message
}

func NewAppError(code int, message string) *AppError {
	return &AppError{code, message}
}

func ErrorHandler(ctx fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*AppError); ok {
		code = e.Code
	}

	return ctx.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})
}
