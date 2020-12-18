package repository

import (
	"github.com/blackhorseya/todo-app/internal/app/entities"
	"github.com/google/wire"
)

// TaskRepo is a repository to task Business
type TaskRepo interface {
	QueryTaskList(limit, offset int32) (tasks []*entities.Task, err error)
	CreateTask(newTask *entities.Task) (task *entities.Task, err error)
	RemoveTask(id string) (count int, err error)
}

// ProviderSet is a repository of task of provider set
var ProviderSet = wire.NewSet(NewImpl)
