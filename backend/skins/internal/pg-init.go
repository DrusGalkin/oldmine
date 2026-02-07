package internal

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func PostgresInit() *sql.DB {
	db, err := sql.Open("postgres", strConnFromEnv())
	if err != nil {
		log.Fatal("Ошибка подключения БД: " + err.Error())
	}

	if err := db.Ping(); err != nil {
		panic("Нет отклика от БД: " + err.Error())
	}

	log.Println("Успешное подключение к БД")

	return db
}

func strConnFromEnv() string {
	if err := godotenv.Load("./.env"); err != nil {
		panic(".env не найден")
	}

	return fmt.Sprintf(
		"password=%s "+
			"host=%s "+
			"port=%s "+
			"user=%s "+
			"dbname=%s "+
			"sslmode=%s",
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
}
