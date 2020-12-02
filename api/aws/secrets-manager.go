package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"go.uber.org/fx"
)

// ProvideSecretsManagerParams for fx
type ProvideSecretsManagerParams struct {
	fx.In
	session *session.Session
}

// ProvideSecretsManager for fx
func ProvideSecretsManager(
	p ProvideSecretsManagerParams,
) *secretsmanager.SecretsManager {
	sm := secretsmanager.New(p.session)
	return sm
}
