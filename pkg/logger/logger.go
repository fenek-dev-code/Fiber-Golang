package logger

import (
	"go-fiber/config"
	"os"

	"github.com/rs/zerolog"
)

func NewLogger(conf *config.LogConfig) *zerolog.Logger {
	// Initialize and return a new logger instance
	zerolog.SetGlobalLevel(zerolog.Level(conf.Level))
	var logger zerolog.Logger
	if conf.Format == "json" {
		logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

	} else {
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	}
	return &logger
}
