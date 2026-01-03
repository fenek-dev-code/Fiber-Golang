package vacancy

import (
	"go-fiber/pkg/tadapter"
	"go-fiber/views/components"
	"time"

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
	form := CreateVacancyForm(c)
	formsErrs := form.IsValid()
	time.Sleep(2 * time.Second)
	if formsErrs.HasAny() {
		comp := components.Notification(formsErrs.Error(), components.NotificationFail)
		return tadapter.Render(c, comp)
	}
	comp := components.Notification("Вакансия созданна!", components.NotificationSuccess)
	return tadapter.Render(c, comp)
}
