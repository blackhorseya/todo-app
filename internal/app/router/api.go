package router

import (
	// import docs for swagger
	_ "github.com/blackhorseya/todo-app/internal/app/docs"
	"github.com/gin-gonic/gin"
)

// RegisterAPI register api group route
func (r *Router) RegisterAPI(app *gin.Engine) {
	api := app.Group("/api")
	{
		api.GET("/readiness", r.HealthHandler.Readiness)
		api.GET("/liveness", r.HealthHandler.Liveness)

		v1 := api.Group("/v1")
		{
			tasks := v1.Group("/tasks")
			{
				tasks.GET("", r.TaskHandler.List)
				tasks.POST("", r.TaskHandler.Create)
				tasks.DELETE("/:id", r.TaskHandler.Remove)
				tasks.PATCH("/:id", r.TaskHandler.ModifyInfo)
			}
		}
	}
}
