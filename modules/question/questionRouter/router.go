package router

import (
	"github.com/gofiber/fiber/v2"
	middlewareHandler "github.com/gunktp20/digital-hubx-be/modules/middleware/middlewareHandler"
	questionHandler "github.com/gunktp20/digital-hubx-be/modules/question/questionHandler"
)

func SetQuestionRoutes(api fiber.Router, questionHttpHandler questionHandler.QuestionHttpHandlerService, middlewareHandler middlewareHandler.MiddlewareHandlerService) {
	routes := api.Group("/question", middlewareHandler.JwtAuthorization)

	routes.Post("/", questionHttpHandler.CreateQuestion)
	routes.Get("/:class_id/class", questionHttpHandler.GetQuestionsByClassID)

}
