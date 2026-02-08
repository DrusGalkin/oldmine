package main

import (
	"fmt"
	"forum/internal/app"
	"forum/internal/config"
	"forum/pkg"
	"forum/pkg/database"
	"github.com/DrusGalkin/libs"
	"time"
)

func main() {
	pkg.MustLoadMkDir(handler.UPLOADS_PATH)

	cfg := config.MustLoad()
	log := libs.LoggerInit(cfg.Env)
	db := database.PostgresInit()

	http := app.Run(db, cfg, log)

	go http.Listen(
		fmt.Sprintf(":%s", cfg.HTTP.Port),
	)

	libs.GracefulShutdown(func() {
		fmt.Println("Завершение работы Forum сервиса")

		if err := http.ShutdownWithTimeout(cfg.HTTP.ShutdownTimeout); err != nil {
			log.Error("Ошибка завершения работы HTTP сервера" + err.Error())
		}

		if err := db.Close(); err != nil {
			log.Error("Ошибка завершения работы PostgreSQL сервера" + err.Error())
		}

		time.Sleep(1 * time.Second)
		fmt.Println("OK!")
	})
}
