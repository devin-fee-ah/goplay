package lib

import (
	"fmt"
	"time"

	ginzap "github.com/gin-contrib/zap"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Routes modules for registration
type Routes interface {
	Register(gin *gin.Engine)
}

// RequestHandler struct
type RequestHandler struct {
	env    *Env
	engine *gin.Engine
	logger *zap.SugaredLogger
	routes []Routes
}

// NewRequestHandler creates a new request handler
func NewRequestHandler(env *Env, logger *zap.SugaredLogger) *RequestHandler {
	return &RequestHandler{
		env:    env,
		logger: logger,
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
	rh.engine.Run(fmt.Sprintf(":%s", rh.env.Port))
}

// SetupAndRun the request handler
func (rh *RequestHandler) SetupAndRun() {
	rh.Setup()
	rh.Run()
}
