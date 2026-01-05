package app

import (
	"auth/internal/repository"
	"auth/internal/transport/http"
	"auth/pkg/database"
	"database/sql"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"os"
	"time"
)

var (
	sessionKey = []byte(os.Getenv("SESSION_KEY"))
)

func HTTPLoad(db *sql.DB, rdb *database.RedisClient, log *zap.Logger, ttl time.Duration) *fiber.App {
	app := fiber.New()

	repo := repository.New(db, log, ttl*10)

	return http.SetupRouters(app, repo, rdb, ttl)
}
