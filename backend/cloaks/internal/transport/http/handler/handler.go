package handler

import (
	"cloaks/internal/repository"
	"cloaks/internal/transport/grpc/client"
	"github.com/gofiber/fiber/v3"
)

const UPLOAD_PATH = "./uploads"
const FORM_FILE_NAME = "cloak"

type Handler interface {
	Save(ctx fiber.Ctx) error
	Get(ctx fiber.Ctx) error
	Delete(ctx fiber.Ctx) error
}

type CloaksHandler struct {
	repo       repository.Repository
	GRPCClient *client.Auth
}

func New(repo repository.Repository, grpcClient *client.Auth) Handler {
	return &CloaksHandler{
		repo:       repo,
		GRPCClient: grpcClient,
	}
}
