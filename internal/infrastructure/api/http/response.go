package httpserver

import "github.com/gofiber/fiber/v2"

type HTTPResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewHTTPResponse(status int, message string, data interface{}) *HTTPResponse {
	return &HTTPResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func NewHTTPResponseCreated(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(&HTTPResponse{
		Status:  fiber.StatusCreated,
		Message: "created",
		Data:    data,
	})
}

func NewHTTPResponseBadRequest(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusBadRequest).JSON(&HTTPResponse{
		Status:  fiber.StatusBadRequest,
		Message: "bad Request",
		Data:    data,
	})
}
