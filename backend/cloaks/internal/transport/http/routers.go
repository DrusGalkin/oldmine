package http

import (
	"cloaks/internal/config"
	"cloaks/internal/transport/grpc/client"
	"cloaks/internal/transport/http/handler"
	"cloaks/internal/transport/http/middleware"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/swagger/v2"
)

func SetupRouters(hd handler.Handler, grpcClient *client.Auth, path string, cfg *config.Config) *fiber.App {
	app := fiber.New()

	app.Use(
		logger.New(),
		//corsState(cfg.Env),
		cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
			AllowCredentials: true,
			ExposeHeaders:    []string{"Set-Cookie"},
			MaxAge:           86400,
		}),
	)

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/uploads/*", static.New(path))
	app.Get("/:id", hd.Get)

	md := middleware.
		NewAuthMiddleware(
			grpcClient,
		)

	auth := app.Use(md.Authenticate())
	{
		auth.Delete("/:id", hd.Delete)
		auth.Post("/", hd.Save)
	}

	return app
}
