package bootstrap

import (
	"context"

	"dfee/api/controllers"
	"dfee/api/lib"
	"dfee/api/repositories"
	"dfee/api/routes"
	"dfee/api/services"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	controllers.Module,
	lib.Module,
	repositories.Module,
	routes.Module,
	services.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler lib.RequestHandler,
	routes routes.Routes,
	env lib.Env,
	logger lib.Logger,
) {

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting Application")
			logger.Zap.Info("---------------------")
			logger.Zap.Info("------- CLEAN -------")
			logger.Zap.Info("---------------------")

			go func() {
				// migrations.Migrate()
				// middlewares.Setup()
				routes.Setup()
				handler.Gin.Run(env.Port)
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Zap.Info("Stopping Application")
			// database.DB.Close()
			return nil
		},
	})
}
