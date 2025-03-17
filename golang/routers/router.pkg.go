package routers

import (
	"behemoth-pkg/golang/config"
	"behemoth-pkg/golang/middlewares"
	"behemoth-pkg/golang/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IRepositories interface {
}

// RouterConfig encapsulates the configuration for the router
type RouterConfig struct {
	engine *gin.Engine
	repos  IRepositories
}

// NewRouter initializes a new router configuration and returns a configured Gin engine
func NewRouter(routerType string, repos IRepositories) *gin.Engine {
	config := &RouterConfig{
		engine: gin.New(),
		repos:  repos,
	}
	if err := config.initialize(routerType); err != nil {
		log.Fatalf("Failed to initialize router: %v", err)
	}
	return config.engine
}

func InitServer(repos IRepositories, specificPort *string) *http.Server {
	appPort := utils.GetEnv("API_PORT", "3000")
	if specificPort != nil {
		appPort = *specificPort
	}
	handler := NewRouter("api", repos)

	return &http.Server{
		Addr:    appPort,
		Handler: handler,
	}
}

// initialize sets up routes, and returns the engine
func (c *RouterConfig) initialize(routerType string) error {
	c.setupMiddleware()
	c.setupUploadFiles()

	// switch routerType {
	// case "api":
	// 	configRoute.RegisterApiRoutes(c.engine, c.repos)
	// default:
	// 	return fmt.Errorf("invalid router type: %s", routerType)
	// }

	return nil
}

// setupUploadFiles sets up the upload file directories for dev
func (c *RouterConfig) setupUploadFiles() {
	if utils.IsDevelopmentEnv() || utils.IsLocalEnv() {
		c.engine.Static("/uploads", "./tmp_app/uploads/")
	}
}

// setupMiddleware attaches global middleware to the Gin engine
func (c *RouterConfig) setupMiddleware() {
	c.engine.Use(config.CorsConfig())
	c.engine.Use(middlewares.IncludeGinCtxIntoCtx())
	c.engine.Use(middlewares.TraceRequest())
	c.engine.Use(middlewares.SetClientLanguage())
}
