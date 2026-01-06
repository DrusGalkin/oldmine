package http

import (
	"auth/internal/config"
	"auth/internal/repository"
	"auth/internal/transport/http/handler"
	"auth/internal/transport/http/middleware"
	"auth/pkg/database"
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/extractors"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/encryptcookie"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/gofiber/swagger/v2"
	"os"
	"time"
)

const COOKIE_NAME_SESSION = "session_id"

func SetupRouters(app *fiber.App, repo repository.Repository, rdb *database.RedisClient, cfg *config.Config) *fiber.App {
	hd := handler.New(repo)

	sessCfg := session.Config{
		Storage:         rdb,
		CookieDomain:    COOKIE_NAME_SESSION,
		Extractor:       extractors.Extractor{},
		IdleTimeout:     cfg.ServerConfig.SessTTl,
		AbsoluteTimeout: cfg.ServerConfig.AbsoluteSessTTl,
		CookieSecure:    true,
		CookieHTTPOnly:  true,
		KeyGenerator: func() string {
			return fmt.Sprintf("sess_%d", time.Now().Unix())
		},
	}

	app.Use(
		logger.New(),
		cors.New(),
		session.New(sessCfg),
		encryptcookie.New(
			encryptcookie.Config{
				Key: base64.StdEncoding.
					EncodeToString(
						[]byte(os.Getenv("COOKIE_SECRET")),
					),
			},
		),
	)

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Post("/login", hd.Login)
	app.Post("/register", hd.Register)
	app.Get("/logout", hd.Logout)
	app.Get("/logout", hd.Profile)

	auth := app.Use(middleware.AuthMiddleware())
	{
		auth.Get("/profile", hd.Profile)
	}

	return app
}
