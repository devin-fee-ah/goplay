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
	var config zap.Config

	if p.Env.Environment == "development" {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config = zap.NewProductionConfig()
	}

	logger, _ := config.Build()

	return logger.Sugar()
}
