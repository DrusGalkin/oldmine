package main

import (
	"auth/internal/app"
	"auth/internal/config"
	"auth/internal/transport/grpc"
	"auth/pkg/database"
	"fmt"
	"libs"
	"os"
	"time"
)

// @title OldMine Auth Api
// @version 1.0
// @description Документация запросов микросервиса Auth
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8123
func main() {
	cfg := config.MustLoadConfig()

	pg := database.PostgresInit()
	rdb := database.RedisInit()

	// ENV указывать в докере
	log := libs.LoggerInit(os.Getenv("ENV"))
	defer log.Sync()

	// HTTP и gRPC
	/* ********************************************** */

	httpApp := app.HTTPLoad(
		pg,
		rdb,
		log,
		cfg,
	)
	go httpApp.Listen(":" + cfg.ServerConfig.Port)

	grpcApp := grpc.New(
		rdb,
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
