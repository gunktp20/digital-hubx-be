package handler

import (
	middlewareUsecase "github.com/gunktp20/digital-hubx-be/modules/middleware/middlewareUsecase"

	"github.com/gofiber/fiber/v2"
)

type (
	MiddlewareHandlerService interface {
		JwtAuthorization(c *fiber.Ctx) error
	}

	middlewareHandler struct {
		middlewareUsecase middlewareUsecase.MiddlewareUsecaseService
	}
)

func NewMiddlewareHttpHandler(usecase middlewareUsecase.MiddlewareUsecaseService) MiddlewareHandlerService {
	return &middlewareHandler{middlewareUsecase: usecase}
}
