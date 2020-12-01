package lib

import "os"

// Env has environment stored
type Env struct {
	DBHost      string
	DBName      string
	DBPassword  string
	DBPort      string
	DBUsername  string
	Environment string
	Port        string
}

// NewEnv creates a new environment
func NewEnv() Env {
	env := Env{}
	env.LoadEnv()
	return env
}

// LoadEnv loads environment
func (env *Env) LoadEnv() {
	env.DBHost = os.Getenv("DBHost")
	env.DBName = os.Getenv("DBName")
	env.DBPassword = os.Getenv("DBPassword")
	env.DBPort = os.Getenv("DBPort")
	env.DBUsername = os.Getenv("DBUsername")
	env.Environment = os.Getenv("Environment")
	env.Port = os.Getenv("Port")
}
