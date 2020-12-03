package ent

import (
	"context"

	cfg "github.com/devin-fee-ah/goplay/config"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module for fx
var Module = fx.Options(
	fx.Provide(ProvideEnt),
	fx.Invoke(bootstrap),
)

// ProvideEntParams for fx
type ProvideEntParams struct {
	fx.In
	Config *cfg.Config
	Logger *zap.SugaredLogger
}

// ProvideEnt for fx
func ProvideEnt(p ProvideEntParams) (client *Client, err error) {
	client, err = Open("sqlite3", p.Config.DatabaseURL)
	if err != nil {
		p.Logger.Fatalf("Failed to connect to Database: %v", err)
	}
	return
}

func bootstrap(
	db *Client,
	logger *zap.SugaredLogger,
	lifecycle fx.Lifecycle,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) (err error) {
			// Run the auto migration tool.
			if err = db.Schema.Create(context.Background()); err != nil {
				logger.Fatalf("Failed to create schema resources: %v", err)
			}
			return
		},
		OnStop: func(context.Context) (err error) {
			// Close the database
			err = db.Close()
			if err != nil {
				logger.Fatalf("Failed to close Database: %v", err)
			}
			return
		},
	})
}
