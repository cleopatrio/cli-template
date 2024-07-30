package html

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// Router - decorates the provided server with HTML-only routes.
func Router(router fiber.Router) {
	router.
		Get("", func(c *fiber.Ctx) error {
			return c.Render(
				"templates/index",
				fiber.Map{
					"time": time.Now(),
				},
				"templates/layouts/base",
			)
		}).
		Name("web:index")
}
