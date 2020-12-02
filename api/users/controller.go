package users

import (
	"dfee/api/lib"
	"dfee/api/users/dtos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ControllerParams for fx
type ControllerParams struct {
	fx.In
	Logger      *zap.SugaredLogger
	UserService *Service
}

// Controller data type
type Controller struct {
	userService *Service
	logger      *zap.SugaredLogger
}

// NewController creates new user controller
func NewController(p ControllerParams) *Controller {
	return &Controller{
		userService: p.UserService,
		logger:      p.Logger,
	}
}

// GetOne gets one user
// @Summary Get a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} ent.User
// @Failure 400 {object} lib.HTTPError
// @Failure 404 {object} lib.HTTPError
// @Failure 500 {object} lib.HTTPError
// @Router /v1/users/{id} [get]
func (c *Controller) GetOne(ctx *gin.Context) {
	paramID := ctx.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		lib.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	user, err := c.userService.GetOne(ctx, id)
	if err != nil {
		lib.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// GetAll gets all users
// @Summary Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} []ent.User
// @Failure 400 {object} lib.HTTPError
// @Failure 404 {object} lib.HTTPError
// @Failure 500 {object} lib.HTTPError
// @Router /v1/users/ [get]
func (c *Controller) GetAll(ctx *gin.Context) {
	users, err := c.userService.GetAll(ctx)
	if err != nil {
		lib.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

// Save saves a new user
// @Summary Saves a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body dtos.AddUser true "Add user"
// @Success 200 {object} ent.User
// @Failure 400 {object} lib.HTTPError
// @Failure 404 {object} lib.HTTPError
// @Failure 500 {object} lib.HTTPError
// @Router /v1/users [post]
func (c *Controller) Save(ctx *gin.Context) {
	var dto dtos.AddUser
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		lib.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	user, err := c.userService.Create(ctx.Request.Context(), dto)
	if err != nil {
		lib.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

// Update updates a user
// @Summary Update a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body dtos.UpdateUser true "Update user"
// @Success 200 {object} ent.User
// @Failure 400 {object} lib.HTTPError
// @Failure 404 {object} lib.HTTPError
// @Failure 500 {object} lib.HTTPError
// @Router /v1/users/{id} [post]
func (c *Controller) Update(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		lib.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	var dto dtos.UpdateUser
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		lib.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	user, err := c.userService.Update(ctx, id, dto)
	if err != nil {
		lib.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

// Delete deletes a user
// @Summary Delete a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 204 {object} ent.User
// @Failure 400 {object} lib.HTTPError
// @Failure 404 {object} lib.HTTPError
// @Failure 500 {object} lib.HTTPError
// @Router /v1/users/{id} [delete]
func (c *Controller) Delete(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		lib.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	err = c.userService.Delete(ctx, id)
	if err != nil {
		lib.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
