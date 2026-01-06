package grpc

import (
	"auth/pkg/database"
	"fmt"
	"google.golang.org/grpc"
	glogger "libs/logger/grpc-logger"
	"libs/proto/generate"
	"net"
)

type Server struct {
	generate.UnimplementedAuthServer
	redis *database.RedisClient
	grpc  *grpc.Server
	host  string
	port  string
}

func New(rdb *database.RedisClient, host, port string) *Server {
	return &Server{
		redis: rdb,
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
