package app

import (
	"github.com/blackhorseya/todo-app/internal/app/middlewares"
	"github.com/blackhorseya/todo-app/internal/app/router"
	"github.com/blackhorseya/todo-app/internal/pkg/config"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewGinEngine init a engine of Gin
func NewGinEngine(r router.IRouter, config *config.Config) *gin.Engine {
	gin.SetMode(config.RunMode)

	app := gin.New()

	// logger
	app.Use(middlewares.LoggerMiddleware())

	// recovery
	app.Use(gin.Recovery())

	// frontend
	app.Use(static.Serve("/", static.LocalFile("./web/build", true)))

	// register route to Gin engine
	_ = r.Register(app)

	// swagger
	if mode := gin.Mode(); mode == gin.DebugMode {
		app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return app
}
