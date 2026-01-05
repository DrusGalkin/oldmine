package grpc

import (
	"auth/pkg/database"
	"fmt"
	"google.golang.org/grpc"
	glogger "libs/logger/grpc-logger"
	"libs/proto/generate"
	"net"
	"time"
)

type Server struct {
	generate.UnimplementedAuthServer
	ttl   time.Duration
	redis *database.RedisClient
	host  string
	port  string
}

func New(rdb *database.RedisClient, host, port string, ttl time.Duration) *Server {
	return &Server{
		redis: rdb,
		ttl:   ttl,
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

	err = grpcServer.Serve(lis)
	if err != nil {
		panic("Ошибка запуска gRPC сервера: " + err.Error())
	}
}

func recoveryRun(server *Server) {
	if r := recover(); r != nil {
		server.host = "localhost"
		fmt.Println("gRPC сервер запущен на localhost")
		server.Run()
	}
}
