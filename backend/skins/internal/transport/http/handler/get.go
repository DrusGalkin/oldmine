package handler

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

func (h *SkinHandler) Get(ctx fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	path, err := h.repo.Get(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err,
			})
		} else {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"path": path,
	})
}
