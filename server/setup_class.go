package server

import (
	"github.com/gofiber/fiber/v2"
	classHandler "github.com/gunktp20/digital-hubx-be/modules/class/classHandler"
	classRepository "github.com/gunktp20/digital-hubx-be/modules/class/classRepository"
	classRouter "github.com/gunktp20/digital-hubx-be/modules/class/classRouter"
	classUsecase "github.com/gunktp20/digital-hubx-be/modules/class/classUsecase"
	classCategoryRepository "github.com/gunktp20/digital-hubx-be/modules/classCategory/classCategoryRepository"
	middlewareHandler "github.com/gunktp20/digital-hubx-be/modules/middleware/middlewareHandler"
	middlewareUsecase "github.com/gunktp20/digital-hubx-be/modules/middleware/middlewareUsecase"
	"github.com/gunktp20/digital-hubx-be/pkg/config"
)

func (s *fiberServer) initializeClassHttpHandler(api fiber.Router, conf *config.Config) {
	// ? Initialize all layers
	classGormRepository := classRepository.NewClassGormRepository(s.db.GetDb())
	classCategoryGormRepository := classCategoryRepository.NewClassCategoryGormRepository(s.db.GetDb())

	classUsecase := classUsecase.NewClassUsecase(classGormRepository, classCategoryGormRepository, s.bucket)
	classHttpHandler := classHandler.NewClassHttpHandler(classUsecase)

	middlewareUsecase := middlewareUsecase.NewMiddlewareUsecase(s.conf)
	middlewareHandler := middlewareHandler.NewMiddlewareHttpHandler(middlewareUsecase)

	// Routers
	classRouter.SetClassRoutes(api, classHttpHandler, middlewareHandler)
}
