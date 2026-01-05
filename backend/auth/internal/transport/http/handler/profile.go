package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

func (h *AuthHandler) Profile(c fiber.Ctx) error {
	sess := session.FromContext(c)

	return c.JSON(fiber.Map{
		"id":    sess.Get("id").(int),
		"email": sess.Get("email").(string),
		"name":  sess.Get("name").(string),
	})
}
