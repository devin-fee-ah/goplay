package users

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Routes struct
type Routes struct {
	logger     *zap.SugaredLogger
	controller *Controller
}

// NewRoutes creates new user routes
func NewRoutes(
	controller *Controller,
	logger *zap.SugaredLogger,
) *Routes {
	return &Routes{
		logger:     logger,
		controller: controller,
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
				users.GET("/:id", r.controller.GetOneUser)
				// v1.GET("/user", s.userController.GetUser)
				// v1.POST("/user", s.userController.SaveUser)
				// v1.POST("/user/:id", s.userController.UpdateUser)
				// v1.DELETE("/user/:id", s.userController.DeleteUser)
			}
		}
	}
}
