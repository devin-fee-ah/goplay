package awsutils

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"go.uber.org/fx"
)

// ProvideSessionParams for fx
type ProvideSessionParams struct {
	fx.In
}

// ProvideSession for fx
func ProvideSession(p ProvideSessionParams) (*session.Session, error) {
	return session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
}
