package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"libs/proto/generate"
	"time"
)

type Auth struct {
	client  generate.AuthClient
	timeout time.Duration
}

func New(host, port string, timeout time.Duration) *Auth {
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	defer conn.Close()
	if err != nil {
		panic("Ошибка подключения к gRPC серверу" + err.Error())
	}

	return &Auth{
		client:  generate.NewAuthClient(conn),
		timeout: timeout,
	}
}

func (a Auth) getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), a.timeout)
}
