package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

func AuthMiddleware() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		sess := session.FromContext(ctx)

		auth := sess.Get("auth")

		if auth == nil || auth.(bool) == false {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Требуется авторизация",
			})
		}

		return ctx.Next()
	}
}
