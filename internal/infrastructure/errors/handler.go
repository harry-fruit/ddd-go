package apperror

import (
	"github.com/gofiber/fiber/v2"
	httpserver "github.com/harry-fruit/ddd-go/internal/infrastructure/api/http"
)

func Handle(c *fiber.Ctx, err error) error {
	if appErr, ok := err.(AppError); ok {
		return c.Status(appErr.StatusCode()).JSON(&httpserver.HTTPResponse{
			Status:  appErr.StatusCode(),
			Message: appErr.Message(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(&httpserver.HTTPResponse{
		Status:  fiber.StatusInternalServerError,
		Message: "internal server error",
		Data:    nil,
	})
}
