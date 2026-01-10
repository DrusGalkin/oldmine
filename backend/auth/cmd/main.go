package main

import (
	_ "auth/docs"
	"auth/internal/app"
	"auth/internal/config"
	"auth/internal/transport/grpc"
	"auth/pkg/database"
	"fmt"
	"libs"
	"os"
	"time"
)

// @title OldMine Auth Api by AndrewGalkin
// @version 1.0
// @description Документация запросов микросервиса Auth, по всем вопросам @andrewandrew05
// @contact.name Почта
// @contact.email drus.galkin@mail.ru
// @securityDefinitions.apikey CookieAuth
// @in cookie
// @name session_id
// @description Session authentication with cookie
func main() {
	cfg := config.MustLoadConfig()

	pg := database.PostgresInit()
	rdb := database.RedisInit()

	// ENV указывать в докере
	log := libs.LoggerInit(os.Getenv("ENV"))
	defer log.Sync()

	// HTTP и gRPC
	/* ********************************************** */

	httpApp, repo := app.HTTPLoad(
		pg,
		rdb,
		log,
		cfg,
	)
	go httpApp.Listen(fmt.Sprintf(":%d", cfg.ServerConfig.Port))

	grpcApp := grpc.New(
		rdb,
		repo,
		cfg.GRPCConfig.Host,
		cfg.GRPCConfig.Port,
	)
	go grpcApp.Run()

	/* ********************************************** */

	libs.GracefulShutdown(func() {
		fmt.Println("Завершение работы Auth сервиса")

		if err := httpApp.ShutdownWithTimeout(cfg.ServerConfig.ShutdownTimeout); err != nil {
			log.Error("Ошибка завершения работы HTTP сервера" + err.Error())
		}

		grpcApp.Stop()

		if err := pg.Close(); err != nil {
			log.Error("Ошибка завершения работы PostgreSQL сервера" + err.Error())
		}

		if err := rdb.Close(); err != nil {
			log.Error("Ошибка завершения работы Redis сервера" + err.Error())
		}

		time.Sleep(1 * time.Second)
		fmt.Println("OK!")
	})

}
