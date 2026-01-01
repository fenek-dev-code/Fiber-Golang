package vacancy

import (
	"go-fiber/pkg/tadapter"
	"go-fiber/views/components"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type VacnacyHanlder struct {
	router fiber.Router
	log    *zerolog.Logger
}

func NewVacancyHanlder(router fiber.Router, log *zerolog.Logger) {
	r := VacnacyHanlder{
		router: router,
		log:    log,
	}
	r.router.Post("/vacancy", r.create)
}

func (h *VacnacyHanlder) create(c *fiber.Ctx) error {
	email := c.FormValue("email")
	h.log.Info().Msg(email)
	comp := components.Notification("Вакансия созданна!")
	return tadapter.Render(c, comp)
}
