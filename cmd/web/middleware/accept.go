package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SupportedMediaTypes(types ...string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		accepts := c.Accepts(types...)

		if accepts != "" {
			logrus.Debugln("Accepting request...", c.Get("Accept"), c.Path(), c.GetReqHeaders())
			return c.Next()
		}

		logrus.Debugln("1 -", c.Accepts(types...))
		logrus.Debugln("2 -", c.Get("Content-Type"))
		logrus.Debugln("3 -", c.Get("Accept"))

		return c.SendStatus(http.StatusUnsupportedMediaType)
	}
}
