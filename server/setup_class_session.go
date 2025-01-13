package server

import (
	"github.com/gofiber/fiber/v2"
	classSessionHandler "github.com/gunktp20/digital-hubx-be/modules/classSession/classSessionHandler"
	classSessionRepository "github.com/gunktp20/digital-hubx-be/modules/classSession/classSessionRepository"
	classSessionRouter "github.com/gunktp20/digital-hubx-be/modules/classSession/classSessionRouter"
	classSessionUsecase "github.com/gunktp20/digital-hubx-be/modules/classSession/classSessionUsecase"
	middlewareHandler "github.com/gunktp20/digital-hubx-be/modules/middleware/middlewareHandler"
	middlewareUsecase "github.com/gunktp20/digital-hubx-be/modules/middleware/middlewareUsecase"
	"github.com/gunktp20/digital-hubx-be/pkg/config"
)

func (s *fiberServer) initializeClassSessionHttpHandler(api fiber.Router, conf *config.Config) {
	// ? Initialize all layers
	classSessionGormRepository := classSessionRepository.NewClassSessionGormRepository(s.db.GetDb())

	classSessionUsecase := classSessionUsecase.NewClassSessionUsecase(classSessionGormRepository, s.bucket)
	classSessionHttpHandler := classSessionHandler.NewClassSessionHttpHandler(classSessionUsecase)

	middlewareUsecase := middlewareUsecase.NewMiddlewareUsecase(s.conf)
	middlewareHandler := middlewareHandler.NewMiddlewareHttpHandler(middlewareUsecase)

	// Routers
	classSessionRouter.SetClassSessionRoutes(api, classSessionHttpHandler, middlewareHandler)
}
