package handler

import (
	"github.com/gofiber/fiber/v3"
	"skins/internal/repository"
	"skins/internal/transport/grpc/client"
)

const UPLOAD_PATH = "./uploads"

type Handler interface {
	Save(ctx fiber.Ctx) error
	Get(ctx fiber.Ctx) error
	Delete(ctx fiber.Ctx) error
	Update(ctx fiber.Ctx) error
}

type SkinHandler struct {
	repo       repository.Repository
	GRPCClient *client.Auth
}

func New(repo repository.Repository, grpcClient *client.Auth) Handler {
	return &SkinHandler{
		repo:       repo,
		GRPCClient: grpcClient,
	}
}
