package main

import (
	"auth/internal/app"
	"auth/internal/transport/grpc"
	"auth/pkg/database"
	"fmt"
	"libs"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	pg := database.PostgresInit()
	rdb := database.RedisInit()

	log := libs.LoggerInit(os.Getenv("ENV"))
	defer log.Sync()

	go app.HTTPLoad(
		pg,
		rdb,
		log,
		4*time.Hour,
	).Listen(":8123")

	go grpc.New(
		rdb,
		"localhost",
		"50051",
		2000,
	).Run()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGHUP)

	select {
	case <-sigs:
		fmt.Println("Завершение работы Auth сервиса")
	}

}
