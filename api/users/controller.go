package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Controller data type
type Controller struct {
	userService *Service
	logger      *zap.SugaredLogger
}

// NewController creates new user controller
func NewController(
	logger *zap.SugaredLogger,
	userService *Service,
) *Controller {
	return &Controller{
		userService: userService,
		logger:      logger,
	}
}

// GetOneUser gets one user
// @Summary Get a user
// @Description get string by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} dtos.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /users/{id} [get]
func (c *Controller) GetOneUser(ctx *gin.Context) {
	paramID := ctx.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	user, err := c.userService.GetOneUser(int(id))

	if err != nil {
		c.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": user,
	})
}
