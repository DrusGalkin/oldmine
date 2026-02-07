package main

import (
	"log"
	"migrator/internal/app"
	"migrator/pkg/database"
	"os"
)

func main() {
	db := database.PostgresInit()

	switch os.Getenv("state") {
	case "up":
		app.Up(db)
	case "down":
		app.Down(db)
	default:
		log.Println("Миграции не выполняются")
	}
}
