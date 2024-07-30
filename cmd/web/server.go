package web

import (
	"context"
	"embed"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	fiberHTML "github.com/gofiber/template/html/v2"

	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/skip"
	"github.com/google/uuid"
	"github.com/oleoneto/redic/app"
	"github.com/oleoneto/redic/cmd/web/middleware"
	"github.com/oleoneto/redic/cmd/web/negotiators/html"
)

//go:embed public
var public embed.FS

//go:embed templates/*
var templates embed.FS

func CreateServer() *fiber.App {
	views := fiberHTML.NewFileSystem(http.FS(templates), ".html")

	server := fiber.New(fiber.Config{
		AppName:               "clitemplate",
		ServerHeader:          "clitemplate",
		DisableStartupMessage: true,
		ReadTimeout:           5 * time.Second,
		PassLocalsToViews:     true,
		Views:                 views,
	})

	server.Use(recover.New(recover.Config{EnableStackTrace: true}))
	server.Use(requestid.New(requestid.Config{Generator: uuid.NewString}))
	server.Use(limiter.New(limiter.Config{Max: 25}))
	server.Use(csrf.New(csrf.Config{SingleUseToken: true}))

	server.Use(favicon.New(favicon.Config{File: "public/favicon.ico", FileSystem: http.FS(public)}))

	server.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			_, err := app.DatabaseEngine.ExecContext(context.TODO(), "SELECT 1")
			return err == nil
		},
		ReadinessProbe: func(c *fiber.Ctx) bool {
			_, err := app.DatabaseEngine.ExecContext(context.TODO(), "SELECT 1")
			return err == nil
		},
	}))

	server.Use(middleware.RequestLoggerMiddleware)

	server.Get("/stack", skip.New(
		func(c *fiber.Ctx) error { return c.JSON(c.App().Stack()) },
		func(c *fiber.Ctx) bool { return !c.IsFromLocal() },
	))

	server.Get("/pprof", skip.New(
		pprof.New(pprof.Config{}),
		func(c *fiber.Ctx) bool { return !c.IsFromLocal() },
	))

	htmx := server.Group("")
	htmx.Route("", html.Router).
		Use(middleware.SupportedMediaTypes("text/html"))

	return server
}
