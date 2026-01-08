package test

import (
	"auth/internal/models"
	"auth/internal/repository"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"libs"
	"testing"
	"time"
)

func TestAuthRepository_Get(t *testing.T) {
	log, _ := zap.NewDevelopment()

	email := "test@mail.com"
	password := "0123120"
	hashPass := libs.HashPass(password)

	tUser := models.User{
		ID:        1,
		Name:      "bigBoy",
		Email:     email,
		Password:  hashPass,
		CreatedAt: time.Now(),
	}

	t.Run("Успешный поиск пользователя", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer db.Close()

		repo := repository.New(db, log, 2*time.Second)

		rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at"}).
			AddRow(
				tUser.ID,
				tUser.Name,
				tUser.Email,
				tUser.Password,
				tUser.CreatedAt,
			)

		mock.
			ExpectQuery(
				`select id, name, email, password, created_at from users where email = \$1`,
			).
			WithArgs(email).
			WillReturnRows(rows)

		user, err := repo.GetUser(email, password)

		assert.NoError(t, err)
		assert.Equal(t, tUser.ID, user.ID)
		assert.Equal(t, tUser.Email, user.Email)
		assert.Equal(t, tUser.Name, user.Name)
		assert.Equal(t, tUser.CreatedAt, user.CreatedAt)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Пользователь не найден", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer db.Close()

		repo := repository.New(db, log, 2*time.Second)

		mock.
			ExpectQuery(
				`select id, name, email, password, created_at from users where email = \$1`,
			).
			WithArgs(email)

		user, err := repo.GetUser(email, password)

		assert.Error(t, err)
		assert.Equal(t, 0, user.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
