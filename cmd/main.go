package main

import (
	"go-fiber/app"
	"go-fiber/config"
	"go-fiber/pkg/database"
	"go-fiber/pkg/logger"
)

func main() {
	config.LoadEnv()
	cfg := config.GetEnvConfig()
	logger := logger.NewLogger(cfg.LogConfig)

	db := database.NewDbPool(cfg.DatabaseURL)
	defer database.CloseDB(db)

	app := app.App(app.AppDeps{
		Logger:    logger,
		EnvConfig: cfg,
		Pool:      db,
	})
	defer app.Shutdown()

	if err := app.Listen(":8081"); err != nil {
		logger.Fatal().Err(err)
		return
	}
}
