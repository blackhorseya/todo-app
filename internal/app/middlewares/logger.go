package middlewares

import (
	"net/http"
	"os"
	"time"

	"github.com/blackhorseya/todo-app/internal/pkg/ctx"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoggerMiddleware using Logrus for gin access log
func LoggerMiddleware() gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "UNKNOWN"
	}

	return func(c *gin.Context) {
		// start timer
		startAt := time.Now()

		// process request
		c.Next()

		// end timer
		duration := time.Now().Sub(startAt)
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		method := c.Request.Method
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}
		traceID := c.MustGet("ctx").(ctx.CTX).Value("traceID").(string)

		entry := logrus.WithFields(logrus.Fields{
			"traceID":    traceID,
			"duration":   duration,
			"hostname":   hostname,
			"statusCode": statusCode,
			"clientIP":   clientIP,
			"method":     method,
			"path":       path,
			"referer":    referer,
			"dataLength": dataLength,
			"userAgent":  clientUserAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			if statusCode >= http.StatusInternalServerError {
				entry.Error()
			} else if statusCode >= http.StatusBadRequest {
				entry.Warn()
			} else {
				entry.Info()
			}
		}
	}
}
