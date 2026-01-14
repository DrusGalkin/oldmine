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

	app.All("/auth/*", func(c fiber.Ctx) error {
		return proxy.Do(c, "http://auth:8123"+c.Params("*"))
	})

	app.All("/skins/*", func(c fiber.Ctx) error {
		return proxy.Do(c, "http://skins:8122"+c.Params("*"))
	})

	app.Listen(":3000")
}
