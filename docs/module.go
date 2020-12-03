package docs

import (
	"dfee/api/lib"

	"go.uber.org/fx"
)

// Module for fx
var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Invoke(func(handler *lib.RequestHandler, routes *Routes) {
		handler.RegisterRoutes(routes)
	}),
)
