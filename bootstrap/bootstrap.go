package bootstrap

import (
	"context"

	"github.com/devin-fee-ah/goplay/awsutils"
	"github.com/devin-fee-ah/goplay/config"
	"github.com/devin-fee-ah/goplay/docs"
	"github.com/devin-fee-ah/goplay/ent"
	"github.com/devin-fee-ah/goplay/lib"
	"github.com/devin-fee-ah/goplay/secrets"
	"github.com/devin-fee-ah/goplay/users"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module exported for initializing application
var Module = fx.Options(
	awsutils.Module,
	config.Module,
	docs.Module,
	ent.Module,
	lib.Module,
	secrets.Module,
	users.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	handler *lib.RequestHandler,
	lifecycle fx.Lifecycle,
	logger *zap.SugaredLogger,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				logger.Info("Starting Application")
				handler.SetupAndRun()
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Info("Stopping Application")
			// database.DB.Close()
			return nil
		},
	})
}
