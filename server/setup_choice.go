package server

import (
	"github.com/gofiber/fiber/v2"
	choiceHandler "github.com/gunktp20/digital-hubx-be/modules/choice/choiceHandler"
	choiceRepository "github.com/gunktp20/digital-hubx-be/modules/choice/choiceRepository"
	choiceRouter "github.com/gunktp20/digital-hubx-be/modules/choice/choiceRouter"
	choiceUsecase "github.com/gunktp20/digital-hubx-be/modules/choice/choiceUsecase"
	middlewareHandler "github.com/gunktp20/digital-hubx-be/modules/middleware/middlewareHandler"
	middlewareUsecase "github.com/gunktp20/digital-hubx-be/modules/middleware/middlewareUsecase"
	questionRepository "github.com/gunktp20/digital-hubx-be/modules/question/questionRepository"
	"github.com/gunktp20/digital-hubx-be/pkg/config"
)

func (s *fiberServer) initializeChoiceHttpHandler(api fiber.Router, conf *config.Config) {
	// ? Initialize all layers
	choiceGormRepository := choiceRepository.NewChoiceGormRepository(s.db.GetDb())
	questionGormRepository := questionRepository.NewQuestionGormRepository(s.db.GetDb())

	choiceUsecase := choiceUsecase.NewChoiceUsecase(choiceGormRepository, questionGormRepository)
	choiceHttpHandler := choiceHandler.NewChoiceHttpHandler(choiceUsecase)

	middlewareUsecase := middlewareUsecase.NewMiddlewareUsecase(s.conf)
	middlewareHandler := middlewareHandler.NewMiddlewareHttpHandler(middlewareUsecase)

	// Routers
	choiceRouter.SetChoiceRoutes(api, choiceHttpHandler, middlewareHandler)
}
