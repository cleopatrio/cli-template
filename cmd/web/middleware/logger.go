package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RequestLoggerMiddleware(c *fiber.Ctx) error {
	s := time.Now()

	err := c.Next()

	e := time.Now()

	queries := []string{}
	for k, v := range c.Queries() {
		queries = append(queries, fmt.Sprintf("%s=%s", k, v))
	}

	url := fmt.Sprintf("%s%s%s%s",
		c.Hostname(),
		c.Path(),
		func() string {
			if len(queries) > 0 {
				return "?"
			}
			return ""
		}(),
		strings.Join(queries, "&"),
	)

	logrus.Infoln(
		c.Protocol(),
		c.Method(),
		url, "-",
		c.Response().StatusCode(),
		"Duration:", e.Sub(s),
	)

	return err
}
