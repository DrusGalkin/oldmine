package test

import (
	"auth/internal/mock"
	mgrpc "auth/internal/transport/grpc"
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"libs/proto/generate"
	"testing"
	"time"
)

func TestCheckAuth_Client(t *testing.T) {
	sessID := "sess_123"
	rdb := mock.NewMockRedis()

	chErr := make(chan error, 1)
	host, port := "localhost", "50051"
	grpcServer := mgrpc.New(
		rdb,
		nil,
		host,
		port,
	)

	sessionData := map[interface{}]interface{}{
		"id":    1,
		"name":  "user",
		"email": "example@example.com",
		"auth":  true,
	}

	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)

	gob.Register(sessionData)

	if err := encoder.Encode(sessionData); err != nil {
		t.Fatal("Ошибка сериализации: " + err.Error())
	}

	data := buf.Bytes()

	err := rdb.Set(sessID, data, 2*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				chErr <- err.(error)
			}
		}()
		grpcServer.Run()
	}()

	time.Sleep(200 * time.Millisecond)

	go func() {
		conn, _ := grpc.NewClient(
			fmt.Sprintf("%s:%s", host, port),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		defer conn.Close()

		client := generate.NewAuthClient(conn)
		req := &generate.AuthRequest{
			SessId: sessID,
		}

		res, err := client.CheckAuth(context.Background(), req)
		if err != nil {
			chErr <- err
		}

		t.Log(res)
	}()

	select {
	case err := <-chErr:
		t.Fatal(err)
	case <-time.After(1 * time.Second):
		t.Log("TestCheckAuth_Client OK")
	}
}
