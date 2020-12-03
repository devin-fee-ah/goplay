package awsutils

import (
	"go.uber.org/fx"
)

// Module for fx
var Module = fx.Options(
	fx.Provide(NewToolbelt),
	fx.Provide(ProvideSecretsManager),
	fx.Provide(ProvideSession),
)
