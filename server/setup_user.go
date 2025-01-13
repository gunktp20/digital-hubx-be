package server

import (
	userHandler "github.com/gunktp20/digital-hubx-be/modules/user/userHandler"
	userRepository "github.com/gunktp20/digital-hubx-be/modules/user/userRepository"
	userRouter "github.com/gunktp20/digital-hubx-be/modules/user/userRouter"
	userUsecase "github.com/gunktp20/digital-hubx-be/modules/user/userUsecase"

	"github.com/gofiber/fiber/v2"
)

func (s *fiberServer) initializeUserHttpHandler(api fiber.Router) {
	// ? Initialize all layers
	userPostgresRepository := userRepository.NewUserGormRepository(s.db.GetDb())

	userUsecase := userUsecase.NewUserUsecase(userPostgresRepository)
	userHttpHandler := userHandler.NewUserHttpHandler(userUsecase)
	// Routers
	userRouter.SetupUserRoutes(api, userHttpHandler)
}
