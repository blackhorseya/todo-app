package task

import (
	"github.com/blackhorseya/todo-app/internal/app/biz/task/repository"
	"github.com/blackhorseya/todo-app/internal/app/entities"
	"github.com/google/wire"
)

// Biz describe task business service function
type Biz interface {
	Create(newTask *entities.Task) (task *entities.Task, err error)
	Remove(id string) (count int, err error)
	UpdateStatus(id string, completed bool) (task *entities.Task, err error)
	ChangeTitle(id, newTitle string) (task *entities.Task, err error)
	List(page, size int32) (tasks []*entities.Task, err error)
}

// ProviderSet is a task provider set
var ProviderSet = wire.NewSet(NewImpl, repository.ProviderSet)
