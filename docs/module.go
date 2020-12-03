package docs

import (
	"github.com/devin-fee-ah/goplay/lib"

	"go.uber.org/fx"
)

// Module for fx
var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Invoke(func(handler *lib.RequestHandler, routes *Routes) {
		handler.RegisterRoutes(routes)
	}),
)
