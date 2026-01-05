package libs

import (
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
