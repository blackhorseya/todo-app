package health

import (
	"net/http"

	"github.com/blackhorseya/todo-app/internal/app/biz/health"
	"github.com/blackhorseya/todo-app/internal/app/entities"
	"github.com/gin-gonic/gin"
)

type impl struct {
	HealthBiz health.Biz
}

// NewImpl is a constructor of implement health api handler
func NewImpl(healthBiz health.Biz) IHandler {
	return &impl{HealthBiz: healthBiz}
}

// Readiness to know when an application is ready to start accepting traffic
// @Summary Readiness
// @Description Show application was ready to start accepting traffic
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /readiness [get]
func (h *impl) Readiness(ctx *gin.Context) {
	res := &entities.Response{
		Ok:  true,
		Msg: "application has been ready",
	}
	code := http.StatusOK

	ok, _ := h.HealthBiz.Readiness()
	if !ok {
		code = http.StatusInternalServerError
		res.Ok = false
		res.Msg = "application has failed"
	}

	ctx.JSON(code, res)
}

// Liveness to know when to restart an application
// @Summary Liveness
// @Description to know when to restart an application
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /liveness [get]
func (h *impl) Liveness(ctx *gin.Context) {
	ret := &entities.Response{
		Ok:  true,
		Msg: "alive",
	}
	code := http.StatusOK

	ok, _ := h.HealthBiz.Liveness()
	if !ok {
		code = http.StatusInternalServerError
		ret.Ok = false
		ret.Msg = "dead"
	}

	ctx.JSON(code, ret)
}
