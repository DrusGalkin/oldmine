package grpc

import (
	"auth/internal/dto"
	"context"
	"database/sql"
	"errors"
	glogger "github.com/DrusGalkin/libs/logger/grpc-logger"
	"github.com/DrusGalkin/libs/proto/generate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) PaymentVerification(
	ctx context.Context,
	req *generate.PaymentVerificationRequest,
) (*generate.PaymentVerificationResponse, error) {
	const op = "grpc.PaymentVerification"

	boolCh := make(chan bool, 1)
	errCh := make(chan dto.GRPCError, 1)

	go func() {
		ver, err := s.repo.PaymentVerification(ctx, req.Id)
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

		boolCh <- ver
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

	return &generate.PaymentVerificationResponse{
		Pay: <-boolCh,
	}, glogger.Print(op, codes.OK)
}
