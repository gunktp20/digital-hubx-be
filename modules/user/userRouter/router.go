package router

import (
	handler "github.com/gunktp20/digital-hubx-be/modules/user/userHandler"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(api fiber.Router, userHttpHandler handler.UserHttpHandlerService) {
	routes := api.Group("/user")

	routes.Post("/", userHttpHandler.CreateUser)
}
