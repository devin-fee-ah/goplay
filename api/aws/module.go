package aws

import "go.uber.org/fx"

// Module for fx
var Module = fx.Options(
	fx.Provide(ProvideSession),
	fx.Provide(ProvideSecretsManager),
)
