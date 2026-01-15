package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/proxy"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(), logger.New())

	api := app.Group("/api")
	{
		api.All("/auth/*", func(c fiber.Ctx) error {
			return proxy.Do(c, "http://auth:8123"+c.Params("*"))
		})

		app.All("/skins/*", func(c fiber.Ctx) error {
			return proxy.Do(c, "http://skins:8122"+c.Params("*"))
		})
	}

	docs := app.Group("/docs")
	{
		docs.Get("/auth/*", func(c fiber.Ctx) error {
			return proxy.Do(c, "http://auth:8123/swagger"+c.Params("*"))
		})
		docs.Get("/skins/*", func(c fiber.Ctx) error {
			return proxy.Do(c, "http://skins:8122/swagger"+c.Params("*"))
		})
	}

	app.Listen(":3000")
}
