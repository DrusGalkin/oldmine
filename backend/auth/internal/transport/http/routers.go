package http

import (
	"auth/internal/config"
	"auth/internal/transport/http/handler"
	"auth/internal/transport/http/middleware"
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

func SetupRouters(app *fiber.App, hd handler.Handler, rdb fiber.Storage, cfg *config.Config) *fiber.App {
	sessCfg := session.Config{
		Storage:         rdb,
		CookieDomain:    COOKIE_NAME_SESSION,
		Extractor:       extractors.Extractor{},
		IdleTimeout:     cfg.ServerConfig.SessTTl,
		AbsoluteTimeout: cfg.ServerConfig.AbsoluteSessTTl,
		CookieSecure:    false,
		CookieHTTPOnly:  true,
		KeyGenerator: func() string {
			return fmt.Sprintf("sess_%d", time.Now().Unix())
		},
	}

	app.Use(
		logger.New(),
		cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
			AllowCredentials: true,
			ExposeHeaders:    []string{"Set-Cookie"},
			MaxAge:           86400,
		}),
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

	auth := app.Use(middleware.AuthMiddleware())
	{
		auth.Get("/profile", hd.Profile)
	}

	return app
}
