package server

import (
	"fmt"

	"github.com/gunktp20/digital-hubx-be/pkg/bucket"
	"github.com/gunktp20/digital-hubx-be/pkg/config"
	"github.com/gunktp20/digital-hubx-be/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/gunktp20/digital-hubx-be/docs"
)

type (
	Server interface {
		Start()
	}

	fiberServer struct {
		app    *fiber.App
		db     database.Database
		conf   *config.Config
		bucket bucket.BucketClientService
	}
)

func NewFiberServer(conf *config.Config, db database.Database, bucket bucket.BucketClientService) Server {
	fiberApp := fiber.New(fiber.Config{
		ReadBufferSize:        60 * 1024,
		DisableStartupMessage: false,
	})

	return &fiberServer{
		app:    fiberApp,
		db:     db,
		conf:   conf,
		bucket: bucket,
	}
}

func (s *fiberServer) Start() {
	s.app.Use(logger.New())
	s.app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173,http://example.com",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	// s.app.Get("/swagger/*", swagger.HandlerDefault)

	api := s.app.Group("/api")
	s.app.Get("", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Good health ✅",
		})
	})

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)

	s.initializeUserHttpHandler(api)
	s.initializeAuthHttpHandler(api, s.conf)
	s.initializeClassHttpHandler(api, s.conf)
	s.initializeClassCategoryHttpHandler(api, s.conf)
	s.initializeClassSessionHttpHandler(api, s.conf)
	s.initializeClassRegistrationHttpHandler(api, s.conf)
	s.initializeQuestionHttpHandler(api, s.conf)
	s.initializeChoiceHttpHandler(api, s.conf)

	s.app.Listen(serverUrl)
}
