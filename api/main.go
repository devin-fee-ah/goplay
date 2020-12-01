package main

import (
	"dfee/api/bootstrap"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/fx"
)

// @title Swagger Example API
// @version 1.0
// @description This is a simple app
// @contact.name Devin Fee
// @contact.email devin.fee@earnin.cocm
// @host localhost:8080
// @BasePath /api/v1
func main() {
	godotenv.Load()
	fx.New(bootstrap.Module).Run()
}
