package router

import (
	authHandler "github.com/gunktp20/digital-hubx-be/modules/auth/authHandler"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(api fiber.Router, authHttpHandler authHandler.AuthHttpHandlerService) {
	routes := api.Group("/auth")

	routes.Post("/login", authHttpHandler.Authenticate)

}
