package router

import (
	"github.com/gofiber/fiber/v2"
	classSessionHandler "github.com/gunktp20/digital-hubx-be/modules/classSession/classSessionHandler"
	middlewareHandler "github.com/gunktp20/digital-hubx-be/modules/middleware/middlewareHandler"
)

func SetClassSessionRoutes(api fiber.Router, classSessionHttpHandler classSessionHandler.ClassSessionHttpHandlerService, middlewareHandler middlewareHandler.MiddlewareHandlerService) {
	routes := api.Group("/class-session", middlewareHandler.JwtAuthorization)

	routes.Get("/", classSessionHttpHandler.GetAllClassSessions)
	routes.Post("/", classSessionHttpHandler.CreateClassSession)

}
