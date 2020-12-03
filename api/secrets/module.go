package secrets

import (
	"dfee/api/awsutils"
	"dfee/api/lib"

	"go.uber.org/fx"
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
}

// NewSecretsParams for fx
type NewSecretsParams struct {
	fx.In
	Env         *lib.Env
	AwsToolbelt *awsutils.Toolbelt
}

// NewSecrets for fx
func NewSecrets(p NewSecretsParams) *Secrets {
	return &Secrets{
		awsToolbelt: p.AwsToolbelt,
	}
}
