package docs

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// RoutesParams for fx
type RoutesParams struct {
	fx.In
	Logger *zap.SugaredLogger
}

// Routes for swagger
type Routes struct {
	logger *zap.SugaredLogger
}

// NewRoutes creates new user routes
func NewRoutes(p RoutesParams) *Routes {
	return &Routes{
		logger: p.Logger,
	}
}

// Register user routes
func (r *Routes) Register(e *gin.Engine) {
	r.logger.Info("Setting up swagger routes")
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
