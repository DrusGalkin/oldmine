package app

import (
	"database/sql"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"skins/internal/config"
	"skins/internal/repository"
	"skins/internal/transport/grpc/client"
	"skins/internal/transport/http"
	"skins/internal/transport/http/handler"
)

func Run(db *sql.DB, log *zap.Logger, cfg *config.Config) *fiber.App {
	repo := repository.New(db, log)

	grpcClient := client.New(
		cfg.GRPC.Host,
		cfg.GRPC.Port,
		cfg.GRPC.Timeout,
	)

	hd := handler.New(repo, grpcClient)
	return http.SetupRouters(hd, grpcClient, handler.UPLOAD_PATH, cfg)
}
