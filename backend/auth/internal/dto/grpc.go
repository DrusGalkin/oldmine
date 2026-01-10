package dto

import "google.golang.org/grpc/codes"

type GRPCError struct {
	Code codes.Code
	Err  error
}
