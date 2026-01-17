package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"path/filepath"
)

func (h *SkinHandler) Save(ctx fiber.Ctx) error {
	formFile, err := ctx.FormFile("skin")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	fileName := fmt.Sprintf("%s/%s", UPLOAD_PATH, formFile.Filename)

	if filepath.Ext(formFile.Filename) != ".png" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Разрешенный формат только png",
		})
	}

	return ctx.SaveFile(formFile, fileName)
}
