package lib

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ProvideLoggerParams for fx
type ProvideLoggerParams struct {
	fx.In
	Env *Env
}

// ProvideLogger to fx
func ProvideLogger(p ProvideLoggerParams) *zap.SugaredLogger {
	config := zap.NewDevelopmentConfig()

	if p.Env.Environment == "development" {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	logger, _ := config.Build()

	return logger.Sugar()
}
