package router

import (
	"github.com/gofiber/fiber/v2"
	classRegistrationHandler "github.com/gunktp20/digital-hubx-be/modules/classRegistration/classRegistrationHandler"
	middlewareHandler "github.com/gunktp20/digital-hubx-be/modules/middleware/middlewareHandler"
)

func SetClassRegistrationRoutes(api fiber.Router, classRegistrationHttpHandler classRegistrationHandler.ClassRegistrationHttpHandlerService, middlewareHandler middlewareHandler.MiddlewareHandlerService) {
	routes := api.Group("/class-registration", middlewareHandler.JwtAuthorization)

	routes.Get("/", classRegistrationHttpHandler.GetUserRegistrations)
	routes.Post("/", classRegistrationHttpHandler.CreateClassRegistration)

}
