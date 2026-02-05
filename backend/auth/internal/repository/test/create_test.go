package test

import (
	"auth/internal/models"
	"auth/internal/repository"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestAuthRepository_Create(t *testing.T) {
	log, _ := zap.NewDevelopment()

	tUser := models.User{
		Name:     "user",
		Email:    "example@example.com",
		Password: "12345678",
	}

	t.Run("Успешное создание пользователя", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err, "error opening mock db")
		defer db.Close()

		repo := repository.New(db, log, 2*time.Second)

		rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
		mock.ExpectQuery(`insert into users`).
			WithArgs(
				tUser.Name,
				tUser.Email,
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
			).WillReturnRows(rows)

		err = repo.Create(tUser)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	/*  Поля стукруты пользователя, должны быть валидны струкрутным тегам
	 	models.User
		- Name min=3 max=16
		- Email должен иметь вид ппп@ттт.ppp
		...
	*/
	t.Run("Неправильное создание пользователя", func(t *testing.T) {

		testingUsers := []models.User{
			{
				Name:     "",
				Email:    "expmai@m",
				Password: "133",
			},
			{
				Name:     "And",
				Email:    "@d.com",
				Password: "+_1",
			},
			{
				Name:     "OK",
				Email:    "expmail.com",
				Password: "1",
			},
		}

		for i, tableCase := range testingUsers {
			t.Run(fmt.Sprintf("Тест %d", i+1), func(t *testing.T) {
				db, _, err := sqlmock.New()
				require.NoError(t, err, "error opening mock db")
				defer db.Close()

				repo := repository.New(db, log, 2*time.Second)

				err = repo.Create(tableCase)
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "Невалидные аргументы")
			})
		}
	})
}
