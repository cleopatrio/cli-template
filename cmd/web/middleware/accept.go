package middleware

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

func SupportedMediaTypes(types ...string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		accepts := c.Accepts(types...)

		if accepts != "" {
			log.Debug("Accepting request...", c.Get("Accept"), c.Path(), c.GetReqHeaders())
			return c.Next()
		}

		log.Debug("1 -", c.Accepts(types...))
		log.Debug("2 -", c.Get("Content-Type"))
		log.Debug("3 -", c.Get("Accept"))

		return c.SendStatus(http.StatusUnsupportedMediaType)
	}
}
