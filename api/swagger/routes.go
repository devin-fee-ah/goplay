package swagger

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// Routes struct
type Routes struct {
	logger *zap.SugaredLogger
}

// NewRoutes creates new user routes
func NewRoutes(
	logger *zap.SugaredLogger,
) *Routes {
	return &Routes{
		logger: logger,
	}
}

// Register user routes
func (r *Routes) Register(e *gin.Engine) {
	r.logger.Info("Setting up swagger routes")
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
