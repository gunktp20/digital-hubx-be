package handler

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gunktp20/digital-hubx-be/pkg/response"
)

func (h *middlewareHandler) JwtAuthorization(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return response.ErrResponse(c, http.StatusUnauthorized, "Authorization header is missing", nil)
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return response.ErrResponse(c, http.StatusUnauthorized, "Invalid authorization format", nil)
	}

	claims, err := h.middlewareUsecase.JwtAuthorization(tokenParts[1])
	if err != nil {
		return response.ErrResponse(c, http.StatusUnauthorized, err.Error(), nil)
	}

	// ดึง user_id จาก claims และเก็บใน context
	if userID, ok := claims["user_id"].(string); ok {
		c.Locals("user_id", userID)
	} else {
		return response.ErrResponse(c, http.StatusUnauthorized, "user_id not found in token claims", nil)
	}

	return c.Next()

}
