package secrets

import (
	"dfee/api/awsutils"
	"dfee/api/lib"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module for fx
var Module = fx.Options(
	fx.Provide(NewSecrets),
)

// Secrets for fx
type Secrets struct {
	awsToolbelt *awsutils.Toolbelt
	env         *lib.Env
	goPlay      *GoPlay
	logger      *zap.SugaredLogger
}

// NewSecretsParams for fx
type NewSecretsParams struct {
	fx.In
	AwsToolbelt *awsutils.Toolbelt
	Env         *lib.Env
	Logger      *zap.SugaredLogger
}

// NewSecrets for fx
func NewSecrets(p NewSecretsParams) *Secrets {
	return &Secrets{
		awsToolbelt: p.AwsToolbelt,
		env:         p.Env,
		logger:      p.Logger,
	}
}
