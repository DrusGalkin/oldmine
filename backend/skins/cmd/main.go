package main

import (
	"fmt"
	"libs"
	"skin_system/internal/app"
	"skin_system/internal/config"
	"skin_system/pkg"
	"skin_system/pkg/database"
	"time"
)

const UPLOAD_PATH = "./uploads"

func main() {
	pkg.MustLoadMkDir(UPLOAD_PATH)
	db := database.PostgresInit()

	cfg := config.MustLoad()
	log := libs.LoggerInit(cfg.Env)

	http := app.Run(db, log, cfg)

	libs.GracefulShutdown(func() {
		fmt.Println("Завершение работы Skins сервиса")

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
