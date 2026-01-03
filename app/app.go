package app

import (
	"go-fiber/config"
	"go-fiber/internal/home"
	"go-fiber/internal/vacancy"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type AppDeps struct {
	*zerolog.Logger
	*config.EnvConfig
	*pgxpool.Pool
}

func App(deps AppDeps) *fiber.App {
	app := fiber.New()

	// Init Middleware
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: deps.Logger,
	}))
	app.Use(recover.New())
	app.Static("/public/", "./public")

	// Init Routes
	home.NewHomeHandler(app, deps.Logger)
	vacancy.NewVacancyHanlder(app, deps.Logger)
	return app
}
