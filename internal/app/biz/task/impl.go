package task

import (
	"errors"
	"fmt"
	"time"

	"github.com/blackhorseya/todo-app/internal/app/biz/task/repository"
	"github.com/blackhorseya/todo-app/internal/app/entities"
	"github.com/google/uuid"
)

type impl struct {
	TaskRepo repository.TaskRepo
}

// NewImpl is a constructor of implement business
func NewImpl(repo repository.TaskRepo) Biz {
	return &impl{
		TaskRepo: repo,
	}
}

// Create serve user to create a task
func (i *impl) Create(newTask *entities.Task) (task *entities.Task, err error) {
	if len(newTask.Title) == 0 {
		return nil, fmt.Errorf("title must be NOT empty")
	}

	newTask.Id = uuid.New().String()
	newTask.CreateAt = time.Now().UnixNano()

	task, err = i.TaskRepo.CreateTask(newTask)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// UpdateStatus serve user to update complete status of task by id
func (i *impl) UpdateStatus(id string, completed bool) (task *entities.Task, err error) {
	panic("implement me")
}

// Remove serve user to remove a task by id
func (i *impl) Remove(id string) (count int, err error) {
	if len(id) == 0 {
		return 0, errors.New("id must be NOT empty")
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return 0, err
	}

	count, err = i.TaskRepo.RemoveTask(uid.String())
	if err != nil {
		return 0, err
	}

	return count, nil
}

// ChangeTitle serve user to update title of task
func (i *impl) ChangeTitle(id, newTitle string) (task *entities.Task, err error) {
	panic("implement me")
}

// List all tasks
func (i *impl) List(page, size int32) (tasks []*entities.Task, err error) {
	tasks, err = i.TaskRepo.QueryTaskList(size, (page-1)*size)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
