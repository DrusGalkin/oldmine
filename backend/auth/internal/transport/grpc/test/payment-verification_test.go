package test

import (
	mock2 "auth/internal/mock"
	"auth/internal/repository"
	grpc2 "auth/internal/transport/grpc"
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"libs/proto/generate"
	"testing"
	"time"
)

func TestPaymentVerification_Client(t *testing.T) {
	log, _ := zap.NewDevelopment()
	var uid int64 = 1

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := repository.New(db, log, 2*time.Second)

	rows := sqlmock.NewRows([]string{"user_id"}).AddRow(1)
	mock.
		ExpectQuery(
			`select user_id from paid_for where user_id = \$1`,
		).
		WithArgs(uid).
		WillReturnRows(rows)

	host, port := "localhost", "50052"

	errCh := make(chan error, 1)
	resCh := make(chan bool, 1)

	server := grpc2.New(new(mock2.Redis), repo, host, port)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				errCh <- err.(error)
			}
			server.Run()
		}()
	}()

	go func() {
		conn, _ := grpc.NewClient(
			fmt.Sprintf("%s:%s", host, port),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)

		client := generate.NewAuthClient(conn)
		ver, err := client.PaymentVerification(context.Background(), &generate.PaymentVerificationRequest{Id: uid})
		if err != nil {
			errCh <- err
		}
		resCh <- ver.Pay
	}()

	select {
	case err := <-errCh:
		require.NoError(t, err)
	case res := <-resCh:
		assert.True(t, res)
	}
}
