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
	// Application entry point

	// Load configuration and initialize logger
	config.LoadEnv()
	cfg := config.GetEnvConfig()

	// Set up logger
	logger := logger.NewLogger(cfg.LogConfig)

	// Initialize Fiber app
	app := fiber.New()

	// Middleware
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: logger,
	}))
	app.Use(recover.New())

	// Register home handler
	home.NewHomeHandler(app, logger)

	// Start server
	_ = app.Listen(":8081")
}
