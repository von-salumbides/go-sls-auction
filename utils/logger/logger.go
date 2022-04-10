package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitZap() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()
}

func ERROR(msg, err string) {
	zap.L().Error(msg, zap.Any("error", err))
}

func INFO(msg string, fields ...zapcore.Field) {
	zap.L().Info(msg, fields...)
}
