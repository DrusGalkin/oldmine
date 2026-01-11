package test

import (
	"auth/internal/mock"
	"auth/internal/transport/grpc"
	"testing"
	"time"
)

func TestGRPCServer(t *testing.T) {
	grpcSever := grpc.New(
		new(mock.Redis),
		nil,
		"localhost",
		"",
	)

	chErr := make(chan error, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				chErr <- r.(error)
			}
		}()

		grpcSever.Run()
	}()

	time.Sleep(1 * time.Second)
	grpcSever.Stop()

	select {
	case err := <-chErr:
		t.Fatal(err)
	case <-time.After(1 * time.Second):
		t.Log("TestGRPCServer OK")
	}
}
