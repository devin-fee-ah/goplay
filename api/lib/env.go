package lib

import (
	"dfee/api/utils"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"go.uber.org/fx"
)

var environmentOptions = []string{"development", "production"}

// EnvParams for fx (placeholder)
type EnvParams struct {
	fx.In
}

// Env has environment stored
type Env struct {
	DatabaseURL string
	Environment string
	Port        uint
}

// NewEnv creates a new environment
func NewEnv(p EnvParams) (env *Env, err error) {
	env = &Env{}
	err = env.Load()
	return
}

// Load environment
func (env *Env) Load() (err error) {
	databaseURL := os.Getenv("DATABASE_URL")
	environment := os.Getenv("ENVIRONMENT")
	port := os.Getenv("PORT")

	if len(databaseURL) == 0 {
		return errors.New("DATABASE_URL must be provided")
	}
	env.DatabaseURL = databaseURL

	if !utils.StringInSlice(environment, environmentOptions) {
		return fmt.Errorf(
			"ENVIRONMENT must be in [%s]",
			strings.Join(environmentOptions, ", "),
		)
	}
	env.Environment = environment

	if len(port) == 0 {
		env.Port = 8080
	} else {
		portUint64, err := strconv.ParseUint(port, 10, 16)
		if err != nil {
			return err
		}
		env.Port = uint(portUint64)
	}

	return nil
}
