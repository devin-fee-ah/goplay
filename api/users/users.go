package users

import (
	"dfee/api/lib"

	"go.uber.org/fx"
)

// Module exports services present
var Module = fx.Options(
	fx.Provide(NewController),
	fx.Provide(NewRepository),
	fx.Provide(NewRoutes),
	fx.Provide(NewService),
	fx.Invoke(func(routes *Routes, handler *lib.RequestHandler) {
		handler.RegisterRoutes(routes)
	}),
)
