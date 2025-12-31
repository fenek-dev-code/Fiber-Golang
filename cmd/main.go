package main

import (
	"go-fiber/config"
	"go-fiber/internal/home"
	"go-fiber/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.LoadEnv()
	cfg := config.GetEnvConfig()
	logger := logger.NewLogger(cfg.LogConfig)

	app := fiber.New()

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: logger,
	}))
	app.Use(recover.New())
	app.Static("/public/", "./public")

	home.NewHomeHandler(app, logger)

	if err := app.Listen(":8081"); err != nil {
		logger.Fatal().Err(err)
		return
	}
}
