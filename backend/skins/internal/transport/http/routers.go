package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/swagger/v2"
	"skins/internal/config"
	"skins/internal/transport/grpc/client"
	"skins/internal/transport/http/handler"
	"skins/internal/transport/http/middleware"
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

	app.Get("/uploads/skins/*", static.New(path))
	app.Get("/:id", hd.Get)

	if cfg.Env == "dev" {

		app.Delete("/:id", hd.Delete)
		app.Put("/", hd.Update)
		app.Post("/", hd.Save)

	} else {
		md := middleware.
			NewAuthMiddleware(
				grpcClient,
			)

		auth := app.Use(md.Authenticate())
		{
			auth.Delete("/:id", hd.Delete)
			auth.Put("/", hd.Update)
			auth.Post("/", hd.Save)
		}
	}

	return app
}

func corsState(env string) fiber.Handler {
	if env == "prod" {
		return cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
			AllowCredentials: true,
			ExposeHeaders:    []string{"Set-Cookie"},
			MaxAge:           86400,
		})
	}
	return cors.New()
}
