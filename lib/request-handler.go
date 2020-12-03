package lib

import (
	"fmt"
	"time"

	ginzap "github.com/gin-contrib/zap"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Routes modules for registration
type Routes interface {
	Register(gin *gin.Engine)
}

// RequestHandlerParams for fx
type RequestHandlerParams struct {
	fx.In
	Env    *Env
	Logger *zap.SugaredLogger
}

// RequestHandler struct
type RequestHandler struct {
	env    *Env
	engine *gin.Engine
	logger *zap.SugaredLogger
	routes []Routes
}

// NewRequestHandler creates a new request handler
func NewRequestHandler(p RequestHandlerParams) *RequestHandler {
	return &RequestHandler{
		env:    p.Env,
		logger: p.Logger,
		routes: []Routes{},
	}
}

// RegisterRoutes and return a new Route
func (rh *RequestHandler) RegisterRoutes(routes Routes) {
	rh.routes = append(rh.routes[:], routes)
}

// Setup engine and register routes
func (rh *RequestHandler) Setup() *gin.Engine {
	if rh.engine == nil {
		rh.engine = gin.New()

		plainLogger := rh.logger.Desugar()
		rh.engine.Use(ginzap.Ginzap(plainLogger, time.RFC3339, true))
		rh.engine.Use(ginzap.RecoveryWithZap(plainLogger, true))

		for _, registration := range rh.routes {
			registration.Register(rh.engine)
		}
	}

	return rh.engine
}

// Run the request handler
func (rh *RequestHandler) Run() {
	rh.engine.Run(fmt.Sprintf(":%d", rh.env.Port))
}

// SetupAndRun the request handler
func (rh *RequestHandler) SetupAndRun() {
	rh.Setup()
	rh.Run()
}
