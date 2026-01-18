package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"os"
	"path/filepath"
)

func (h *SkinHandler) Update(ctx fiber.Ctx) error {
	formFile, err := ctx.FormFile(FORM_FILE_NAME)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	if filepath.Ext(formFile.Filename) != ".png" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Разрешенный формат только png",
		})
	}

	id := ctx.Locals("id").(int)
	fileName := ctx.Locals("name").(string) + ".png"
	filePath := fmt.Sprintf(
		"%s/%s",
		UPLOAD_PATH,
		fileName,
	)

	if err := h.repo.Update(ctx.Context(), id, filePath); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	os.Remove(filePath)
	if err := ctx.SaveFile(formFile, filePath); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{})
}
