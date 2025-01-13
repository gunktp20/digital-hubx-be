package response

import "github.com/gofiber/fiber/v2"

type (
	MsgResponse struct {
		Message string             `json:"message"`
		Status  int                `json:"status"`
		Details *map[string]string `json:"details"`
	}
)

func ErrResponse(c *fiber.Ctx, statusCode int, message string, details *map[string]string) error {
	return c.Status(statusCode).JSON(&MsgResponse{Status: statusCode, Message: message, Details: details})
}

func SuccessResponse(c *fiber.Ctx, statusCode int, data any) error {
	return c.Status(statusCode).JSON(data)
}
