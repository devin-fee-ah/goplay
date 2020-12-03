package awsutils

// Use this code snippet in your app.
// If you need more information about configurations or implementing the sample code, visit the AWS docs:
// https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/setting-up.html

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"go.uber.org/fx"
)

// Toolbelt wrapper for fx
type Toolbelt struct {
	secretsManager *secretsmanager.SecretsManager
	session        *session.Session
}

// ToolbeltParams for fx
type ToolbeltParams struct {
	fx.In
	SecretsManager *secretsmanager.SecretsManager
	Session        *session.Session
}

// NewToolbelt builder
func NewToolbelt(p ToolbeltParams) *Toolbelt {
	return &Toolbelt{
		secretsManager: p.SecretsManager,
		session:        p.Session,
	}
}
