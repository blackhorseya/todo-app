package task

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/todo-app/internal/app/biz/task"
	"github.com/blackhorseya/todo-app/internal/app/entities"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"google.golang.org/protobuf/types/known/anypb"
)

type impl struct {
	TaskBiz task.Biz
}

// NewImpl is a constructor of implement task api handler
func NewImpl(taskBiz task.Biz) IHandler {
	return &impl{TaskBiz: taskBiz}
}

// List all tasks
// @Summary List
// @Description list all tasks
// @Tags Task
// @Accept application/json
// @Produce application/json
// @Param page query integer false "page" default(1)
// @Param size query integer false "size of page" default(3)
// @Success 200 {array} entities.Task
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/tasks [get]
func (i *impl) List(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePublic)
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", "3"))
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePublic)
	}

	tasks, _ := i.TaskBiz.List(int32(page), int32(size))
	if len(tasks) == 0 {
		c.String(http.StatusNoContent, "")
	}

	c.JSON(http.StatusOK, tasks)
}

// Create a task
// @Summary Create
// @Description create a task
// @Tags Task
// @Accept application/json
// @Produce application/json
// @Param newTask body entities.Task false "new task"
// @Success 200 {object} entities.Task
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/tasks [post]
func (i *impl) Create(c *gin.Context) {
	var newTask *entities.Task
	if err := c.ShouldBindBodyWith(&newTask, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, &entities.Response{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}
	if newTask == nil {
		c.JSON(http.StatusBadRequest, &entities.Response{
			Ok:  false,
			Msg: "missing new task",
		})
		return
	}

	data, err := i.TaskBiz.Create(newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &entities.Response{
			Ok:  false,
			Msg: err.Error(),
		})
		return
	}

	any, _ := anypb.New(data)
	ret := &entities.Response{
		Ok:   true,
		Data: any,
	}

	c.JSON(http.StatusCreated, ret)
}
