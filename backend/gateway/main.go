package main

import (
	"fmt"
	"net/url"
	"strings"

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
		api.All("/auth/*", createProxy("http://localhost:8123"))
		api.All("/skins/*", createProxy("http://localhost:8122"))
	}

	docs := app.Group("/docs")
	{
		docs.All("/auth/*", createSwaggerProxy("http://localhost:8123"))
		docs.All("/skins/*", createSwaggerProxy("http://localhost:8122"))
	}

	app.Listen(":4000")
}

func createProxy(baseURL string) fiber.Handler {
	return func(c fiber.Ctx) error {
		path := c.Params("*")

		queryString := buildQueryString(c.Queries())

		targetURL := baseURL + "/" + path
		if queryString != "" {
			targetURL += "?" + queryString
		}

		return proxy.Do(c, targetURL)
	}
}

func createSwaggerProxy(baseURL string) fiber.Handler {
	return func(c fiber.Ctx) error {
		path := c.Params("*")

		queryString := buildQueryString(c.Queries())

		var targetURL string

		if path == "" || path == "/" {
			targetURL = baseURL + "/swagger/index.html"
		} else if strings.Contains(path, "swagger.json") ||
			strings.Contains(path, "openapi.json") ||
			strings.Contains(path, "swagger.yaml") ||
			strings.Contains(path, "openapi.yaml") {
			targetURL = baseURL + "/" + path
		} else {
			targetURL = baseURL + "/swagger/" + path
		}

		if queryString != "" {
			targetURL += "?" + queryString
		}

		c.Set("Access-Control-Allow-Origin", "*")

		return proxy.Do(c, targetURL)
	}
}

func buildQueryString(queries map[string]string) string {
	if len(queries) == 0 {
		return ""
	}

	var queryParts []string
	for key, value := range queries {
		encodedKey := url.QueryEscape(key)
		encodedValue := url.QueryEscape(value)
		queryParts = append(queryParts, fmt.Sprintf("%s=%s", encodedKey, encodedValue))
	}

	return strings.Join(queryParts, "&")
}
