package app

import (
	"database/sql"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"skin_system/internal/config"
	"skin_system/internal/repository"
	"skin_system/internal/transport/grpc/client"
	"skin_system/internal/transport/http"
)

func Run(db *sql.DB, log *zap.Logger, cfg *config.Config) *fiber.App {
	repo := repository.New(db, log)

	grpcClient := client.New(
		cfg.GRPC.Host,
		cfg.GRPC.Port,
		cfg.GRPC.Timeout,
	)

	handler := handler.New(repo, grpcClient)
	return http.SetupRouters()
}
