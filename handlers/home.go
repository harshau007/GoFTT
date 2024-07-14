package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harshau007/gotempl/components"
	"github.com/harshau007/gotempl/services"
)

func Home(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return services.Render(c, components.Home())
	})
}
