package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

func (h *AuthHandler) Logout(c fiber.Ctx) error {
	sess := session.FromContext(c)

	if err := sess.Destroy(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ошибка удалении сессии",
		})
	}

	c.ClearCookie("session_id")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Выход из аккаунта",
	})
}
