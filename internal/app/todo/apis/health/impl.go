package health

import (
	"net/http"

	"github.com/blackhorseya/todo-app/internal/app/todo/biz/health"
	"github.com/blackhorseya/todo-app/internal/pkg/base/contextx"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/response"
	"github.com/gin-gonic/gin"
)

type impl struct {
	biz health.IBiz
}

// NewImpl is a constructor of implement health api handler
func NewImpl(biz health.IBiz) IHandler {
	return &impl{biz: biz}
}

// Readiness to know when an application is ready to start accepting traffic
// @Summary Readiness
// @Description Show application was ready to start accepting traffic
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response
// @Failure 500 {object} er.APPError
// @Router /readiness [get]
func (i *impl) Readiness(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	_, err := i.biz.Readiness(ctx)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK)
}

// Liveness to know when to restart an application
// @Summary Liveness
// @Description to know when to restart an application
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response
// @Failure 500 {object} er.APPError
// @Router /liveness [get]
func (i *impl) Liveness(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	_, err := i.biz.Liveness(ctx)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK)
}
