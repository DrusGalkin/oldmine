package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

func AuthMiddleware() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		
		sess := session.FromContext(ctx)

		if sess == nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Сессия не найдена",
			})
		}

		if sess.Fresh() {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Требуется авторизация",
			})
		}

		auth, ok := sess.Get("auth").(bool)
		if !ok || auth == false {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Требуется авторизация",
			})
		}

		if sess.Get("id") == nil || sess.Get("email") == nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Неполные данные сессии",
			})
		}

		return ctx.Next()
	}
}
