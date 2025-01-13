package server

import (
	"github.com/gofiber/fiber/v2"
	classCategoryHandler "github.com/gunktp20/digital-hubx-be/modules/classCategory/classCategoryHandler"
	classCategoryRepository "github.com/gunktp20/digital-hubx-be/modules/classCategory/classCategoryRepository"
	classCategoryRouter "github.com/gunktp20/digital-hubx-be/modules/classCategory/classCategoryRouter"
	classCategoryUsecase "github.com/gunktp20/digital-hubx-be/modules/classCategory/classCategoryUsecase"
	middlewareHandler "github.com/gunktp20/digital-hubx-be/modules/middleware/middlewareHandler"
	middlewareUsecase "github.com/gunktp20/digital-hubx-be/modules/middleware/middlewareUsecase"
	"github.com/gunktp20/digital-hubx-be/pkg/config"
)

func (s *fiberServer) initializeClassCategoryHttpHandler(api fiber.Router, conf *config.Config) {
	// ? Initialize all layers
	classCategoryGormRepository := classCategoryRepository.NewClassCategoryGormRepository(s.db.GetDb())

	classCategoryUsecase := classCategoryUsecase.NewClassCategoryUsecase(classCategoryGormRepository)
	classCategoryHttpHandler := classCategoryHandler.NewClassCategoryHttpHandler(classCategoryUsecase)

	middlewareUsecase := middlewareUsecase.NewMiddlewareUsecase(s.conf)
	middlewareHandler := middlewareHandler.NewMiddlewareHttpHandler(middlewareUsecase)

	// Routers
	classCategoryRouter.SetClassCategoryRoutes(api, classCategoryHttpHandler, middlewareHandler)
}
