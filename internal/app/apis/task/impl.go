package task

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/todo-app/internal/app/biz/task"
	"github.com/blackhorseya/todo-app/internal/app/entities"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type impl struct {
	TaskBiz task.Biz
}

// NewImpl is a constructor of implement task api handler
func NewImpl(taskBiz task.Biz) IHandler {
	return &impl{TaskBiz: taskBiz}
}

type modifyReq struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type modifyBody struct {
	Completed int    `form:"completed" json:"completed"`
	Title     string `form:"title" json:"title"`
}

// ModifyInfo partial update information of task
// @Summary ModifyInfo
// @Description modify information of task
// @Tags Task
// @Accept application/json
// @Produce application/json
// @Param id path string true "Task ID"
// @Param completed query string false "completed of task"
// @Param title query string false "title of task"
// @Success 200 {object} string
// @Success 204 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /v1/tasks/{id} [patch]
func (i *impl) ModifyInfo(c *gin.Context) {
	var req modifyReq
	err := c.ShouldBindUri(&req)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	var body modifyBody
	err = c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if body.Completed != 0 {
		completed := false
		if body.Completed == 2 {
			completed = true
		}
		res, err := i.TaskBiz.UpdateStatus(req.ID, completed)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, res)
		return
	}

	if len(body.Title) != 0 {
		res, err := i.TaskBiz.ChangeTitle(req.ID, body.Title)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, res)
		return
	}

	c.JSON(http.StatusBadRequest, "missing some fields")
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
		return
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", "3"))
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePublic)
		return
	}

	tasks, _ := i.TaskBiz.List(int32(page), int32(size))
	if len(tasks) == 0 {
		c.String(http.StatusNotFound, "")
		return
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

	c.JSON(http.StatusCreated, data)
}

type removeReq struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// Remove a task
// @Summary Remove
// @Description remove a task
// @Tags Task
// @Accept application/json
// @Produce application/json
// @Param id path string true "Task ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Success 404 {object} string
// @Failure 500 {object} string
// @Router /v1/tasks/{id} [delete]
func (i *impl) Remove(c *gin.Context) {
	var req removeReq
	err := c.ShouldBindUri(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	_, err = i.TaskBiz.Remove(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, req.ID)
}
