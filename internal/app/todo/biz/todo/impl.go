package todo

import (
	"time"

	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo/repo"
	"github.com/blackhorseya/todo-app/internal/pkg/base/contextx"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/er"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	repo   repo.IRepo
	node   *snowflake.Node
}

// NewImpl serve caller to create an IBiz
func NewImpl(logger *zap.Logger, repo repo.IRepo, node *snowflake.Node) IBiz {
	return &impl{
		logger: logger.With(zap.String("type", "TodoBiz")),
		repo:   repo,
		node:   node,
	}
}

func (i *impl) GetByID(ctx contextx.Contextx, id int64) (task *todo.Task, err error) {
	ret, err := i.repo.GetByID(ctx, id)
	if err != nil {
		i.logger.Error(er.ErrGetTask.Error(), zap.Error(err), zap.Int64("id", id))
		return nil, er.ErrGetTask
	}
	if ret == nil {
		i.logger.Error(er.ErrTaskNotExists.Error(), zap.Int64("id", id))
		return nil, er.ErrTaskNotExists
	}

	return ret, nil
}

func (i *impl) List(ctx contextx.Contextx, start, end int) (tasks []*todo.Task, total int, err error) {
	if start < 0 {
		i.logger.Error(er.ErrInvalidStart.Error(), zap.Int("start", start), zap.Int("end", end))
		return nil, 0, er.ErrInvalidStart
	}

	if end < 0 {
		i.logger.Error(er.ErrInvalidEnd.Error(), zap.Int("start", start), zap.Int("end", end))
		return nil, 0, er.ErrInvalidEnd
	}

	ret, err := i.repo.List(ctx, end-start+1, start)
	if err != nil {
		i.logger.Error(er.ErrListTasks.Error(), zap.Error(err), zap.Int("start", start), zap.Int("end", end))
		return nil, 0, er.ErrListTasks
	}
	if ret == nil {
		i.logger.Error(er.ErrTaskNotExists.Error(), zap.Int("start", start), zap.Int("end", end))
		return nil, 0, er.ErrTaskNotExists
	}

	total, err = i.repo.Count(ctx)
	if err != nil {
		i.logger.Error(er.ErrCountTask.Error(), zap.Error(err))
		return nil, 0, er.ErrCountTask
	}

	return ret, total, nil
}

func (i *impl) Create(ctx contextx.Contextx, title string) (task *todo.Task, err error) {
	if len(title) == 0 {
		i.logger.Error(er.ErrMissingTitle.Error())
		return nil, er.ErrMissingTitle
	}

	newTask := &todo.Task{
		Id:       i.node.Generate().Int64(),
		Title:    title,
		CreateAt: time.Now().UnixNano(),
	}
	ret, err := i.repo.Create(ctx, newTask)
	if err != nil {
		i.logger.Error(er.ErrCreateTask.Error(), zap.Error(err), zap.String("title", title))
		return nil, er.ErrCreateTask
	}

	return ret, nil
}

func (i *impl) UpdateStatus(ctx contextx.Contextx, id int64, status bool) (task *todo.Task, err error) {
	exists, err := i.repo.GetByID(ctx, id)
	if err != nil {
		i.logger.Error(er.ErrGetTask.Error(), zap.Error(err), zap.Int64("id", id))
		return nil, er.ErrGetTask
	}
	if exists == nil {
		i.logger.Error(er.ErrTaskNotExists.Error(), zap.Int64("id", id))
		return nil, er.ErrTaskNotExists
	}

	exists.Completed = status
	ret, err := i.repo.Update(ctx, exists)
	if err != nil {
		i.logger.Error(er.ErrUpdateStatusTask.Error(), zap.Error(err), zap.Int64("id", id), zap.Bool("status", status))
		return nil, er.ErrUpdateStatusTask
	}

	return ret, nil
}

func (i *impl) ChangeTitle(ctx contextx.Contextx, id int64, title string) (task *todo.Task, err error) {
	if len(title) == 0 {
		i.logger.Error(er.ErrMissingTitle.Error())
		return nil, er.ErrMissingTitle
	}

	exists, err := i.repo.GetByID(ctx, id)
	if err != nil {
		i.logger.Error(er.ErrGetTask.Error(), zap.Error(err), zap.Int64("id", id))
		return nil, er.ErrGetTask
	}
	if exists == nil {
		i.logger.Error(er.ErrTaskNotExists.Error(), zap.Int64("id", id))
		return nil, er.ErrTaskNotExists
	}

	exists.Title = title
	ret, err := i.repo.Update(ctx, exists)
	if err != nil {
		i.logger.Error(er.ErrChangeTitleTask.Error(), zap.Error(err), zap.Int64("id", id), zap.String("title", title))
		return nil, er.ErrChangeTitleTask
	}

	return ret, nil
}

func (i *impl) Delete(ctx contextx.Contextx, id int64) error {
	err := i.repo.Remove(ctx, id)
	if err != nil {
		i.logger.Error(er.ErrDeleteTask.Error(), zap.Error(err), zap.Int64("id", id))
		return er.ErrDeleteTask
	}

	return nil
}
