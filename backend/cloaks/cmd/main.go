package main

import (
	_ "cloaks/docs"
	"cloaks/internal/app"
	"cloaks/internal/config"
	"cloaks/internal/transport/http/handler"
	"cloaks/pkg"
	"cloaks/pkg/database"
	"fmt"
	"github.com/DrusGalkin/libs"
	"time"
)

// @title OldMine Skin Api by AndrewGalkin
// @version 1.0
// @description Документация запросов микросервиса Skin, по всем вопросам @andrewandrew05. Почти для всех запросов требуется авторизация.
// @contact.name Почта
// @contact.email drus.galkin@mail.ru
// @securityDefinitions.apikey CookieAuth
// @in cookie
// @name session_id
// @description Session authentication with cookie
// @host localhost:4000
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
		fmt.Println("Завершение работы Cloaks сервиса")

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
