package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func GetUserIDFromContext(c *fiber.Ctx) (string, error) {
	userID := c.Locals("user_id")
	if userID == nil {
		return "", errors.New("user_id not found in context")
	}

	userIDStr, ok := userID.(string)
	if !ok {
		return "", errors.New("user_id is not a valid string")
	}

	return userIDStr, nil
}
