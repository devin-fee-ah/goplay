package users

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// RoutesParams for fx
type RoutesParams struct {
	fx.In
	Logger         *zap.SugaredLogger
	UserController *Controller
}

// Routes struct
type Routes struct {
	logger     *zap.SugaredLogger
	controller *Controller
}

// NewRoutes creates new user routes
func NewRoutes(p RoutesParams) *Routes {
	return &Routes{
		logger:     p.Logger,
		controller: p.UserController,
	}
}

// Register user routes
func (r *Routes) Register(e *gin.Engine) {
	r.logger.Info("Setting up user routes")
	api := e.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			users := v1.Group("/users")
			{
				users.GET("/", r.controller.GetAll)
				users.POST("/", r.controller.Save)
				users.GET("/:id", r.controller.GetOne)
				users.POST("/:id", r.controller.Update)
				users.DELETE("/:id", r.controller.Delete)
			}
		}
	}
}
