package utils

import (
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func GetLogger() *zap.SugaredLogger {
	if log == nil {
		log = initLogger()
	}
	return log
}

func initLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return logger.Sugar()
}
