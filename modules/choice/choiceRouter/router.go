package router

import (
	"github.com/gofiber/fiber/v2"
	choiceHandler "github.com/gunktp20/digital-hubx-be/modules/choice/choiceHandler"
	middlewareHandler "github.com/gunktp20/digital-hubx-be/modules/middleware/middlewareHandler"
)

func SetChoiceRoutes(api fiber.Router, choiceHttpHandler choiceHandler.ChoiceHttpHandlerService, middlewareHandler middlewareHandler.MiddlewareHandlerService) {
	routes := api.Group("/choice", middlewareHandler.JwtAuthorization)

	routes.Post("/", choiceHttpHandler.CreateChoice)

}
