package server

import (
	"github.com/gofiber/fiber/v2"
	classRepository "github.com/gunktp20/digital-hubx-be/modules/class/classRepository"
	middlewareHandler "github.com/gunktp20/digital-hubx-be/modules/middleware/middlewareHandler"
	middlewareUsecase "github.com/gunktp20/digital-hubx-be/modules/middleware/middlewareUsecase"
	questionHandler "github.com/gunktp20/digital-hubx-be/modules/question/questionHandler"
	questionRepository "github.com/gunktp20/digital-hubx-be/modules/question/questionRepository"
	questionRouter "github.com/gunktp20/digital-hubx-be/modules/question/questionRouter"
	questionUsecase "github.com/gunktp20/digital-hubx-be/modules/question/questionUsecase"
	"github.com/gunktp20/digital-hubx-be/pkg/config"
)

func (s *fiberServer) initializeQuestionHttpHandler(api fiber.Router, conf *config.Config) {
	// ? Initialize all layers
	questionGormRepository := questionRepository.NewQuestionGormRepository(s.db.GetDb())
	classGormRepository := classRepository.NewClassGormRepository(s.db.GetDb())

	questionUsecase := questionUsecase.NewQuestionUsecase(questionGormRepository, classGormRepository)
	questionHttpHandler := questionHandler.NewQuestionHttpHandler(questionUsecase)

	middlewareUsecase := middlewareUsecase.NewMiddlewareUsecase(s.conf)
	middlewareHandler := middlewareHandler.NewMiddlewareHttpHandler(middlewareUsecase)

	// Routers
	questionRouter.SetQuestionRoutes(api, questionHttpHandler, middlewareHandler)
}
