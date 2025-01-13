package server

import (
	"github.com/gofiber/fiber/v2"
	classRepository "github.com/gunktp20/digital-hubx-be/modules/class/classRepository"
	classRegistrationHandler "github.com/gunktp20/digital-hubx-be/modules/classRegistration/classRegistrationHandler"
	classRegistrationRepository "github.com/gunktp20/digital-hubx-be/modules/classRegistration/classRegistrationRepository"
	classRegistrationRouter "github.com/gunktp20/digital-hubx-be/modules/classRegistration/classRegistrationRouter"
	classRegistrationUsecase "github.com/gunktp20/digital-hubx-be/modules/classRegistration/classRegistrationUsecase"
	classSessionRepository "github.com/gunktp20/digital-hubx-be/modules/classSession/classSessionRepository"
	middlewareHandler "github.com/gunktp20/digital-hubx-be/modules/middleware/middlewareHandler"
	middlewareUsecase "github.com/gunktp20/digital-hubx-be/modules/middleware/middlewareUsecase"
	"github.com/gunktp20/digital-hubx-be/pkg/config"
)

func (s *fiberServer) initializeClassRegistrationHttpHandler(api fiber.Router, conf *config.Config) {
	// ? Initialize all layers
	classRegistrationGormRepository := classRegistrationRepository.NewClassRegistrationGormRepository(s.db.GetDb())
	classGormRepository := classRepository.NewClassGormRepository(s.db.GetDb())
	classSessionGormRepository := classSessionRepository.NewClassSessionGormRepository(s.db.GetDb())

	classRegistrationUsecase := classRegistrationUsecase.NewClassRegistrationUsecase(classRegistrationGormRepository, classSessionGormRepository, classGormRepository)
	classRegistrationHttpHandler := classRegistrationHandler.NewClassRegistrationHttpHandler(classRegistrationUsecase)

	middlewareUsecase := middlewareUsecase.NewMiddlewareUsecase(s.conf)
	middlewareHandler := middlewareHandler.NewMiddlewareHttpHandler(middlewareUsecase)

	// Routers
	classRegistrationRouter.SetClassRegistrationRoutes(api, classRegistrationHttpHandler, middlewareHandler)
}
