package web

import (
	"fmt"
	"time"

	"github.com/devin-fee-ah/goplay/config"
	ginzap "github.com/gin-contrib/zap"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Routes modules for registration
type Routes interface {
	Register(gin *gin.Engine)
}

// HandlerParams for fx
type HandlerParams struct {
	fx.In
	Config *config.Config
	Logger *zap.SugaredLogger
}

// Handler struct
type Handler struct {
	config *config.Config
	engine *gin.Engine
	logger *zap.SugaredLogger
	routes []Routes
}

// NewHandler creates a new request handler
func NewHandler(p HandlerParams) *Handler {
	return &Handler{
		config: p.Config,
		logger: p.Logger,
		routes: []Routes{},
	}
}

// RegisterRoutes and return a new Route
func (h *Handler) RegisterRoutes(routes Routes) {
	h.routes = append(h.routes[:], routes)
}

// Setup engine and register routes
func (h *Handler) Setup() *gin.Engine {
	if h.engine == nil {
		if h.config.Environment == "production" {
			gin.SetMode(gin.ReleaseMode)
		}

		h.engine = gin.New()

		plainLogger := h.logger.Desugar()
		if h.config.Environment == "development" {
			h.engine.Use(ginzap.Ginzap(plainLogger, time.RFC3339, true))
		}
		h.engine.Use(ginzap.RecoveryWithZap(plainLogger, true))

		for _, registration := range h.routes {
			registration.Register(h.engine)
		}
	}

	return h.engine
}

// Run the request handler
func (h *Handler) Run() {
	h.engine.Run(fmt.Sprintf(":%d", h.config.Port))
}

// SetupAndRun the request handler
func (h *Handler) SetupAndRun() {
	h.Setup()
	h.Run()
}
