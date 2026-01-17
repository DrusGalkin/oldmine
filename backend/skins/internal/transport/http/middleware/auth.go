package middleware

import (
	"github.com/gofiber/fiber/v3"
	"skins/internal/transport/grpc/client"
)

type AuthMiddleware struct {
	grpcClient *client.Auth
}

func NewAuthMiddleware(grpcClient *client.Auth) *AuthMiddleware {
	return &AuthMiddleware{
		grpcClient: grpcClient,
	}
}

func (a *AuthMiddleware) Authenticate() fiber.Handler {
	return func(c fiber.Ctx) error {
		sess := c.Cookies("session_id")

		if len(sess) == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Неавторизован",
			})
		}

		res := a.grpcClient.CheckAuth(sess)

		if res == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "(Auth) Неавторизован",
			})
		}

		res.Pay = a.grpcClient.PaymentVerification(res.ID)
		res.Admin = a.grpcClient.IsAdmin(res.ID)

		c.Locals("id", res.ID)
		c.Locals("email", res.Email)
		c.Locals("auth", res.Auth)
		c.Locals("name", res.Name)
		c.Locals("pay", res.Pay)
		c.Locals("admin", res.Admin)

		return c.Next()
	}
}
