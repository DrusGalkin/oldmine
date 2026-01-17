package main

import (
	"fmt"
	"libs"
	"skins/internal/app"
	"skins/internal/config"
	"skins/internal/transport/http/handler"
	"skins/pkg"
	"skins/pkg/database"
	"time"
)

// @title OldMine Skin Api by AndrewGalkin
// @version 1.0
// @description Документация запросов микросервиса Skin, по всем вопросам @andrewandrew05
// @contact.name Почта
// @contact.email drus.galkin@mail.ru
// @securityDefinitions.apikey CookieAuth
// @in cookie
// @name session_id
// @description Session authentication with cookie
func main() {
	pkg.MustLoadMkDir(handler.UPLOAD_PATH)
	db := database.PostgresInit()

	cfg := config.MustLoad()

	log := libs.LoggerInit(cfg.Env)
	defer log.Sync()

	http := app.Run(db, log, cfg)

	go http.Listen(
		fmt.Sprintf(":%s", cfg.HTTP.Port),
	)

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
