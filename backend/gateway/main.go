package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/proxy"
)

const (
	SKINS_PROD = "http://skins:8122"
	SKINS_DEV  = "http://localhost:8122"

	AUTH_PROD = "http://auth:8123"
	AUTH_DEV  = "http://localhost:8123"

	CLOAKS_PROD = "http://cloaks:8121"
	CLOAKS_DEV  = "http://localhost:8121"
)

func main() {
	authURL := AUTH_DEV
	skinsURL := SKINS_DEV
	cloaksURL := CLOAKS_DEV

	app := fiber.New()

	app.Use(
		logger.New(),
		cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
			AllowCredentials: true,
			ExposeHeaders:    []string{"Set-Cookie"},
			MaxAge:           86400,
		}),
	)

	// Документация
	// **************************************************************
	app.Get("/skins/swagger/*", func(c fiber.Ctx) error {
		path := c.Params("*")
		return proxy.Do(c, skinsURL+"/swagger/"+path)
	})

	app.Get("/auth/swagger/*", func(c fiber.Ctx) error {
		path := c.Params("*")
		return proxy.Do(c, authURL+"/swagger/"+path)
	})

	app.Get("/swagger/auth/doc.json", func(c fiber.Ctx) error {
		return proxy.Do(c, authURL+"/swagger/doc.json")
	})

	app.Get("/swagger/skins/doc.json", func(c fiber.Ctx) error {
		return proxy.Do(c, skinsURL+"/swagger/doc.json")
	})
	// **************************************************************

	// Все остальные запросы для микросервисов
	// - Auth
	// - Skins
	// **************************************************************
	app.All("/api/auth/*", func(c fiber.Ctx) error {
		path := c.Params("*")

		targetURL := authURL + "/" + path

		proxy.Do(c, targetURL)

		c.Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Set("Access-Control-Allow-Credentials", "true")

		return nil
	})

	app.All("/api/skins/*", func(c fiber.Ctx) error {
		path := c.Params("*")

		targetURL := skinsURL + "/" + path
		proxy.Do(c, targetURL)

		c.Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Set("Access-Control-Allow-Credentials", "true")

		return nil
	})

	app.All("/api/cloaks/*", func(c fiber.Ctx) error {
		path := c.Params("*")

		targetURL := cloaksURL + "/" + path
		proxy.Do(c, targetURL)

		c.Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Set("Access-Control-Allow-Credentials", "true")

		return nil
	})
	// **************************************************************

	// Роуты для системы скинов
	// **************************************************************
	app.Get("/MinecraftSkins/*", func(c fiber.Ctx) error {
		path := c.Params("*")
		return proxy.Do(c, skinsURL+"/uploads"+"/"+path)
	})

	app.Get("/MinecraftCloaks/*", func(c fiber.Ctx) error {
		path := c.Params("*")
		return proxy.Do(c, cloaksURL+"/uploads"+"/"+path)
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
