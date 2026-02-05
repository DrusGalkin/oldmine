package grpc

import (
	"auth/internal/repository"
	"fmt"
	glogger "github.com/DrusGalkin/libs/logger/grpc-logger"
	"github.com/DrusGalkin/libs/proto/generate"
	"github.com/gofiber/fiber/v3"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	generate.UnimplementedAuthServer
	redis fiber.Storage
	grpc  *grpc.Server
	repo  repository.Repository
	host  string
	port  string
}

func New(rdb fiber.Storage, repo repository.Repository, host, port string) *Server {
	return &Server{
		redis: rdb,
		repo:  repo,
		host:  host,
		port:  port,
	}
}

func (s *Server) Run() {
	defer recoveryRun(s)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.host, s.port))
	if err != nil {
		panic("Ошибка запуска gRPC сервера: " + err.Error())
	}

	grpcServer := grpc.NewServer()
	generate.RegisterAuthServer(grpcServer, s)

	glogger.PrintStart(s.host, s.port)
	s.grpc = grpcServer

	err = grpcServer.Serve(lis)
	if err != nil {
		panic("Ошибка запуска gRPC сервера: " + err.Error())
	}
}

func (s *Server) Stop() {
	s.grpc.GracefulStop()
}

func recoveryRun(server *Server) {
	if r := recover(); r != nil {
		server.host = "localhost"
		fmt.Println("gRPC сервер запущен на localhost")
		server.Run()
	}
}
