package libs

import (
	"fmt"
	"go.uber.org/zap"
)

func LoggerInit(env string) *zap.Logger {
	var log *zap.Logger

	if env == "prod" {
		log, _ = zap.NewProduction()
	} else {
		log, _ = zap.NewDevelopment()
	}

	log.Info("Логгер инициализирован")

	return log
}

func QueryError(log *zap.Logger, op string, err error) error {
	msg := fmt.Sprintf("%s: %s", op, err.Error())
	log.Error(msg)
	return err
}
