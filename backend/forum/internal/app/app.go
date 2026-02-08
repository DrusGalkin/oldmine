package app

import (
	"database/sql"
	"forum/internal/config"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func Run(db *sql.DB, cfg *config.Config, log *zap.Logger) *fiber.App {

	return nil
}
