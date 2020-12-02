package bootstrap

import (
	"context"

	"dfee/api/docs"
	"dfee/api/ent"
	"dfee/api/lib"
	"dfee/api/users"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module exported for initializing application
var Module = fx.Options(
	docs.Module,
	ent.Module,
	lib.Module,
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
				// migrations.Migrate()
				// middlewares.Setup()
				// userRoutes.Setup()
				// registrations.Setup()
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
