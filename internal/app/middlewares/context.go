package middlewares

import (
	"github.com/blackhorseya/todo-app/internal/pkg/ctx"
	"github.com/blackhorseya/todo-app/internal/pkg/utils/trace"
	"github.com/gin-gonic/gin"
)

// ContextMiddleware added context into gin
func ContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		context := ctx.WithValue(ctx.Background(), "traceID", trace.NewTraceID())
		c.Set("ctx", context)

		// process request
		c.Next()
	}
}
