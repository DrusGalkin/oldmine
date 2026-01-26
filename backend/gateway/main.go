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

	// Документация
	// **************************************************************
	app.Get("/skins/swagger/*", func(c fiber.Ctx) error {
		path := c.Params("*")
		return proxy.Do(c, "http://skins:8122/swagger/"+path)
	})

	app.Get("/auth/swagger/*", func(c fiber.Ctx) error {
		path := c.Params("*")
		return proxy.Do(c, "http://auth:8123/swagger/"+path)
	})

	app.Get("/swagger/auth/doc.json", func(c fiber.Ctx) error {
		return proxy.Do(c, "http://auth:8123/swagger/doc.json")
	})

	app.Get("/swagger/skins/doc.json", func(c fiber.Ctx) error {
		return proxy.Do(c, "http://skins:8122/swagger/doc.json")
	})
	// **************************************************************

	// Все остальные запросы для микросервисов
	// - Auth
	// - Skins
	// **************************************************************
	app.All("/api/auth/*", func(c fiber.Ctx) error {
		path := c.Params("*")
		return proxy.Do(c, "http://auth:8123/"+path)
	})

	app.All("/api/skins/*", func(c fiber.Ctx) error {
		path := c.Params("*")
		return proxy.Do(c, "http://skins:8122/"+path)
	})
	// **************************************************************

	app.Get("/", func(c fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"message": "Gateway for OldMine!",
			"descriptions": "Чтобы достучаться до документации запросов, надо в строке поиска свагера ввести ссылки " +
				"на docs.json определенного микросервиса. Берешь к примеру ссылку http://localhost:4000/skins/swagger/ или " +
				"http://localhost:4000/auth/swagger/, " +
				"и в поиске указываешь /swagger/skins/doc.json или /swagger/auth/doc.json",

			"endpoints": fiber.Map{
				"api_auth":      "/api/auth/*",
				"api_skins":     "/api/skins/*",
				"swagger_auth":  "/swagger/auth/doc.json",
				"swagger_skins": "/swagger/skins/doc.json",
			},
		})
	})

	app.Listen(":4000")
}
