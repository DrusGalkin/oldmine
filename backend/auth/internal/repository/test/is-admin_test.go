package test

import (
	"auth/internal/repository"
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestIsAdmin(t *testing.T) {
	log, _ := zap.NewDevelopment()
	var uid int64 = 1

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := repository.New(db, log, 2*time.Second)

	rows := sqlmock.NewRows([]string{"user_id"}).AddRow(1)
	mock.
		ExpectQuery(
			`select user_id from admins where user_id = \$1`,
		).
		WithArgs(uid).
		WillReturnRows(rows)

	ver, err := repo.IsAdmin(context.Background(), uid)
	assert.NoError(t, err)
	assert.Equal(t, ver, true)
}
