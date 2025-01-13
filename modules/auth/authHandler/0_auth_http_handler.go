package handler

import (
	"github.com/gofiber/fiber/v2"
	authUsecase "github.com/gunktp20/digital-hubx-be/modules/auth/authUsecase"
)

type (
	AuthHttpHandlerService interface {
		Authenticate(c *fiber.Ctx) error
	}

	authHttpHandler struct {
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthHttpHandler(usecase authUsecase.AuthUsecaseService) AuthHttpHandlerService {
	return &authHttpHandler{authUsecase: usecase}
}
