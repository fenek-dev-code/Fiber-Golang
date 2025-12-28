package home

import (
	"go-fiber/pkg/tadapter"
	"go-fiber/views"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type homeHanlder struct {
	router fiber.Router
	log    *zerolog.Logger
}

func NewHomeHandler(router fiber.Router, log *zerolog.Logger) {
	h := &homeHanlder{router: router, log: log}
	api := h.router.Group("/api")
	api.Get("/", h.handleHome)
	api.Get("/error", h.handleError)
}

func (h *homeHanlder) handleHome(c *fiber.Ctx) error {
	component := views.Main()
	return tadapter.Render(c, component)
}

func (h *homeHanlder) handleError(c *fiber.Ctx) error {
	h.log.Error().Msg("An error occurred in handleError")
	if c != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}
	return fiber.ErrBadRequest
}
