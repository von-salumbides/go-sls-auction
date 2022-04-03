package logger

import "go.uber.org/zap"

func InitZap() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()
}
