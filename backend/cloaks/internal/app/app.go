package app

import (
	"cloaks/internal/config"
	"cloaks/internal/repository"
	"cloaks/internal/transport/grpc/client"
	"cloaks/internal/transport/http"
	"cloaks/internal/transport/http/handler"
	"database/sql"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
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
