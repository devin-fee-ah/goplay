package lib

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ProvideLogger to fx
func ProvideLogger(env *Env) *zap.SugaredLogger {
	config := zap.NewDevelopmentConfig()

	if env.Environment == "development" {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	logger, _ := config.Build()

	return logger.Sugar()
}
