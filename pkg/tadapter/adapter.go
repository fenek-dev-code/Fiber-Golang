package tadapter

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func Render(c *fiber.Ctx, component templ.Component) error {
	// Placeholder for rendering logic
	return adaptor.HTTPHandler(templ.Handler(component))(c)
}
