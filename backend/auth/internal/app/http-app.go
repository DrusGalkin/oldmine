package app

import (
	"auth/internal/config"
	"auth/internal/repository"
	"auth/internal/transport/http"
	"auth/internal/transport/http/handler"
	"database/sql"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func HTTPLoad(db *sql.DB, rdb fiber.Storage, log *zap.Logger, cfg *config.Config) (*fiber.App, repository.Repository) {
	app := fiber.New()

	repo := repository.New(
		db,
		log,
		cfg.ServerConfig.Timeout,
	)

	hd := handler.New(repo)

	return http.SetupRouters(
		app,
		hd,
		rdb,
		cfg,
	), repo
}
