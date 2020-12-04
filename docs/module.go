package docs

import (
	"github.com/devin-fee-ah/goplay/web"
	"go.uber.org/fx"
)

// Module for fx
var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Invoke(func(handler *web.Handler, routes *Routes) {
		handler.RegisterRoutes(routes)
	}),
)
