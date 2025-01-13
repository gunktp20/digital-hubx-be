package server

import (
	authHandler "github.com/gunktp20/digital-hubx-be/modules/auth/authHandler"
	authRouter "github.com/gunktp20/digital-hubx-be/modules/auth/authRouter"
	authUsecase "github.com/gunktp20/digital-hubx-be/modules/auth/authUsecase"
	"github.com/gunktp20/digital-hubx-be/pkg/config"

	userRepository "github.com/gunktp20/digital-hubx-be/modules/user/userRepository"

	"github.com/gofiber/fiber/v2"
)

func (s *fiberServer) initializeAuthHttpHandler(api fiber.Router, conf *config.Config) {
	// ? Initialize all layers
	userPostgresRepository := userRepository.NewUserGormRepository(s.db.GetDb())

	authUsecase := authUsecase.NewAuthUsecase(userPostgresRepository, conf)
	authHttpHandler := authHandler.NewAuthHttpHandler(authUsecase)

	// Routers
	authRouter.SetupAuthRoutes(api, authHttpHandler)
}
