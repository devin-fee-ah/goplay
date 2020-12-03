package config

import (
	"dfee/api/lib"
	"dfee/api/secrets"

	"go.uber.org/fx"
)

// Config for codebase
type Config struct {
	DatabaseURL string
	Environment string
	Port        uint
}

// NewConfigParams for fx
type NewConfigParams struct {
	fx.In
	Env     *lib.Env
	Secrets *secrets.Secrets
}

// NewConfig for fx
func NewConfig(p NewConfigParams) (config *Config, err error) {
	var goPlay *secrets.GoPlay

	databaseURL := p.Env.DatabaseURL
	if len(databaseURL) == 0 {
		goPlay, err = p.Secrets.GetGoPlay()
		if err != nil {
			return
		}
		databaseURL = goPlay.DatabaseURL
	}

	config = &Config{
		DatabaseURL: databaseURL,
		Environment: p.Env.Environment,
		Port:        p.Env.Port,
	}

	return
}
