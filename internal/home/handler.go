package home

import (
	"go-fiber/internal/vacancy"
	"go-fiber/pkg/tadapter"
	"go-fiber/views"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type homeHanlder struct {
	router      fiber.Router
	log         *zerolog.Logger
	vacancyRepo *vacancy.Repository
}

func NewHomeHandler(router fiber.Router, log *zerolog.Logger, vacancyRepo *vacancy.Repository) {
	h := &homeHanlder{router: router, log: log, vacancyRepo: vacancyRepo}
	h.router.Get("/", h.handleHome)
	h.router.Get("/login", h.login)
	h.router.Get("/error", h.handleError)
}

func (h *homeHanlder) handleHome(c *fiber.Ctx) error {
	limit := 2
	page := c.QueryInt("page", 1)
	offset := (page - 1) * limit
	count, err := h.vacancyRepo.Count()
	if err != nil {
		h.log.Error().Err(err).Msg("Failed to get vacancy count")
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get vacancy count")
	}
	vacancies, err := h.vacancyRepo.GetAll(limit, offset)
	if err != nil {
		h.log.Error().Err(err).Msg("Failed to get vacancies")
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get vacancies")
	}

	pagesCount := (count + limit - 1) / limit
	currentPage := page

	component := views.Main(vacancies, pagesCount, currentPage)

	c.Locals("email", "")

	return tadapter.Render(c, component, fiber.StatusOK)
}

func (h *homeHanlder) handleError(c *fiber.Ctx) error {
	h.log.Error().Msg("An error occurred in handleError")
	if c != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}
	return fiber.ErrBadRequest
}

func (h *homeHanlder) login(c *fiber.Ctx) error {
	// Implement login logic here
	return fiber.ErrBadRequest
}
