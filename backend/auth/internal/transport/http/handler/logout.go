package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

// Logout godoc
// @Summary      Выход из аккаунта
// @Description  Завершение сессии и удаление cookies
// @Tags 		 Auth
// @Accept       json
// @Produce      json
// @Security	 CookieAuth
// @Success      200  {object}  object
// @Success		 500  {object}  dto.ErrorResponse
// @Router       /logout [get]
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
