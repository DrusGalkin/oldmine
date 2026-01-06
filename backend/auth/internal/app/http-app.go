package app

import (
	"auth/internal/config"
	"auth/internal/repository"
	"auth/internal/transport/http"
	"auth/pkg/database"
	"database/sql"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func HTTPLoad(db *sql.DB, rdb *database.RedisClient, log *zap.Logger, cfg *config.Config) *fiber.App {
	app := fiber.New()

	repo := repository.New(
		db,
		log,
		cfg.ServerConfig.Timeout,
	)

	return http.SetupRouters(
		app,
		repo,
		rdb,
		cfg,
	)
}
