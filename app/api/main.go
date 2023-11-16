package api

import (
	"chitchat/app"
	"chitchat/app/api/handler"
	"chitchat/config"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

func configure(ac *app.Container) *fiber.App {
	apps := fiber.New()

	apps.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.JSON("PONG")
	})

	apps = loadRoute(apps, ac)

	return apps
}

func Serve(ctx context.Context, ac *app.Container) {
	apps := configure(ac)

	if err := apps.Listen(fmt.Sprintf(":%d", config.Cfg.GetInt("APP_PORT"))); err != nil {
		log.Println("Failed to start server", err.Error())
	}

	fmt.Println("Running cleanup tasks...")
}

func loadRoute(app *fiber.App, ac *app.Container) *fiber.App {
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(pprof.New())

	app = handler.LoadHandler(app, ac)

	return app
}
