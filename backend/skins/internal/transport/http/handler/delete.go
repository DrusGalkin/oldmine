package handler

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v3"
	"os"
	"strconv"
)

func (h *SkinHandler) Delete(ctx fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := h.repo.Delete(ctx.Context(), id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

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

	if err := os.Remove(path); err != nil {
		return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"message": "Ошибка удаления скина, возможно его не существует",
		})
	}

	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Текущий скин удален",
	})
}
