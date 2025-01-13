package router

import (
	"github.com/gofiber/fiber/v2"
	classHandler "github.com/gunktp20/digital-hubx-be/modules/class/classHandler"
	middlewareHandler "github.com/gunktp20/digital-hubx-be/modules/middleware/middlewareHandler"
)

func SetClassRoutes(api fiber.Router, classHttpHandler classHandler.ClassHttpHandlerService, middlewareHandler middlewareHandler.MiddlewareHandlerService) {
	routes := api.Group("/class", middlewareHandler.JwtAuthorization)

	routes.Get("/", classHttpHandler.GetAllClasses)
	routes.Post("/", classHttpHandler.CreateClass)
	routes.Get("/:class_id", classHttpHandler.GetClassById)

}
