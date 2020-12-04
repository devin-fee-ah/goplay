package users

import (
	"github.com/devin-fee-ah/goplay/web"
	"go.uber.org/fx"
)

// Module exports services present
var Module = fx.Options(
	fx.Provide(NewController),
	fx.Provide(NewRoutes),
	fx.Provide(NewService),
	fx.Invoke(func(routes *Routes, handler *web.Handler) {
		handler.RegisterRoutes(routes)
	}),
)
