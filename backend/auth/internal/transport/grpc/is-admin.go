package grpc

import (
	"auth/internal/dto"
	"context"
	"database/sql"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	glogger "libs/logger/grpc-logger"
	"libs/proto/generate"
)

func (s *Server) IsAdmin(
	ctx context.Context,
	req *generate.IsAdminRequest,
) (*generate.IsAdminResponse, error) {
	const op = "grpc.IsAdmin"

	boolCh := make(chan bool, 1)
	errCh := make(chan dto.GRPCError, 1)

	go func() {
		res, err := s.repo.IsAdmin(ctx, req.Id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				errCh <- dto.GRPCError{
					Code: codes.NotFound,
					Err:  err,
				}
			} else {
				errCh <- dto.GRPCError{
					Code: codes.Internal,
					Err:  err,
				}
			}
		}
		boolCh <- res
	}()

	select {
	case <-ctx.Done():
		switch ctx.Err() {
		case context.Canceled:
			return nil, status.Error(codes.Canceled, ctx.Err().Error())
		case context.DeadlineExceeded:
			return nil, status.Error(codes.DeadlineExceeded, ctx.Err().Error())
		}
	case str := <-errCh:
		return nil, glogger.PrintError(op, str.Code, str.Err)
	default:
	}

	return &generate.IsAdminResponse{
		IsAdmin: <-boolCh,
	}, nil
}
