package vacancy

import (
	"go-fiber/pkg/tadapter"
	"go-fiber/views/components"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type Dependencies struct {
	fiber.Router
	*zerolog.Logger
	*Repository
}

type VacnacyHanlder struct {
	router fiber.Router
	log    *zerolog.Logger
	repo   *Repository
}

func NewVacancyHanlder(deps Dependencies) {
	r := VacnacyHanlder{
		router: deps.Router,
		log:    deps.Logger,
		repo:   deps.Repository,
	}
	r.router.Post("/vacancy", r.create)
}

func (h *VacnacyHanlder) create(c *fiber.Ctx) error {
	form := CreateVacancyForm(c)
	formsErrs := form.IsValid()
	if formsErrs.HasAny() {
		comp := components.Notification("Заполните все поля!", components.NotificationFail)
		return tadapter.Render(c, comp, fiber.StatusBadRequest)
	}

	err := h.repo.CreateVacancy(form)
	if err != nil {
		h.log.Error().Err(err).Msg("failed to create vacancy")
		comp := components.Notification("Ошибка создания вакансии", components.NotificationFail)
		return tadapter.Render(c, comp, fiber.StatusInternalServerError)
	}

	comp := components.Notification("Вакансия созданна!", components.NotificationSuccess)
	return tadapter.Render(c, comp, fiber.StatusOK)
}
