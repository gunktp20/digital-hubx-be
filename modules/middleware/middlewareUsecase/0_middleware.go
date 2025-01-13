package usecase

import (
	"github.com/gunktp20/digital-hubx-be/pkg/config"
)

type (
	MiddlewareUsecaseService interface {
		JwtAuthorization(tokenString string) (map[string]interface{}, error)
	}

	middlewareUsecase struct {
		cfg *config.Config
	}
)

func NewMiddlewareUsecase(cfg *config.Config) MiddlewareUsecaseService {
	return &middlewareUsecase{
		cfg: cfg,
	}
}
