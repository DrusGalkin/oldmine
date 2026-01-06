package grpc

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	glogger "libs/logger/grpc-logger"
	"libs/proto/generate"
)

func (s *Server) CheckAuth(ctx context.Context, req *generate.AuthRequest) (*generate.AuthResponse, error) {
	const op = "grpc.CheckAuth"
	data, err := s.redis.GetWithContext(ctx, req.SessId)
	if err != nil {
		return nil, glogger.PrintError(op, codes.Internal, errors.New("Redis error: "+err.Error()))
	}

	if len(data) == 0 {
		return nil, glogger.PrintError(op, codes.NotFound, errors.New("Сессия не найдена"))
	}

	var sessionData map[interface{}]interface{}
	decoder := gob.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&sessionData); err != nil {
		return nil, fmt.Errorf("Ошибка декодирования сессии: %w", err)
	}

	id, _ := sessionData["id"].(int)
	name, _ := sessionData["name"].(string)
	email, _ := sessionData["email"].(string)
	auth, _ := sessionData["auth"].(bool)

	return &generate.AuthResponse{
		Id:    int64(id),
		Name:  name,
		Email: email,
		Auth:  auth,
	}, glogger.Print(op, codes.OK)
}
