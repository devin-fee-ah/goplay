package routes

import (
	"dfee/api/controllers"
	"dfee/api/lib"
)

// UserRoutes struct
type UserRoutes struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	userController controllers.UserController
}

// Setup user routes
func (s UserRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/user/:id", s.userController.GetOneUser)
			// v1.GET("/user", s.userController.GetUser)
			// v1.POST("/user", s.userController.SaveUser)
			// v1.POST("/user/:id", s.userController.UpdateUser)
			// v1.DELETE("/user/:id", s.userController.DeleteUser)
		}
	}
}

// NewUserRoutes creates new user controller
func NewUserRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	userController controllers.UserController,
) UserRoutes {
	return UserRoutes{
		handler:        handler,
		logger:         logger,
		userController: userController,
	}
}
