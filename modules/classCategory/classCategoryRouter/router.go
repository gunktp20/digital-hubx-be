package router

import (
	"github.com/gofiber/fiber/v2"
	classCategoryHandler "github.com/gunktp20/digital-hubx-be/modules/classCategory/classCategoryHandler"
	middlewareHandler "github.com/gunktp20/digital-hubx-be/modules/middleware/middlewareHandler"
)

func SetClassCategoryRoutes(api fiber.Router, classCategoryHttpHandler classCategoryHandler.ClassCategoryHttpHandlerService, middlewareHandler middlewareHandler.MiddlewareHandlerService) {
	routes := api.Group("/class-category", middlewareHandler.JwtAuthorization)

	routes.Get("/", classCategoryHttpHandler.GetAllClassCategories)
	routes.Post("/", classCategoryHttpHandler.CreateClassCategory)

}
