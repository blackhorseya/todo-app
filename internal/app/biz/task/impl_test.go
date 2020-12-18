package task

import (
	"errors"
	"fmt"
	"testing"

	"github.com/blackhorseya/todo-app/internal/app/biz/task/repository/mocks"
	"github.com/blackhorseya/todo-app/internal/app/entities"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type bizTestSuite struct {
	suite.Suite
	taskRepo *mocks.TaskRepo
	taskBiz  Biz
}

func (s *bizTestSuite) SetupSuite() {
}

func (s *bizTestSuite) SetupTest() {
	s.taskRepo = new(mocks.TaskRepo)
	biz, err := CreateTaskBiz(s.taskRepo)
	if err != nil {
		panic(err)
	}
	s.taskBiz = biz
}

func (s *bizTestSuite) TearDownTest() {
	s.taskRepo.AssertExpectations(s.T())
}

func (s *bizTestSuite) Test_impl_Create() {
	type args struct {
		t *entities.Task
	}
	tests := []struct {
		name     string
		args     args
		wantTask *entities.Task
		wantErr  string
		mockFunc func()
	}{
		{
			name: "missing title then nil true",
			args: args{&entities.Task{
				Title: "",
			}},
			wantTask: nil,
			wantErr:  "title must be NOT empty",
			mockFunc: func() {},
		},
		{
			name: "task then task false",
			args: args{&entities.Task{
				Title: "123",
			}},
			wantTask: &entities.Task{
				Title: "123",
			},
			wantErr: "",
			mockFunc: func() {
				s.taskRepo.On("CreateTask", mock.AnythingOfType("*entities.Task")).Return(&entities.Task{
					Title: "123",
				}, nil).Once()
			},
		},
		{
			name: "task then nil true",
			args: args{&entities.Task{
				Title: "456",
			}},
			wantTask: nil,
			wantErr:  "test error",
			mockFunc: func() {
				s.taskRepo.On("CreateTask", mock.AnythingOfType("*entities.Task")).Return(
					nil, fmt.Errorf("test error")).Once()
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mockFunc()
			task, err := s.taskBiz.Create(tt.args.t)
			if err != nil {
				s.EqualErrorf(err, tt.wantErr, "Create() error = [%v], wantErr [%v]", err, tt.wantErr)
			}
			if task != nil {
				s.EqualValuesf(tt.wantTask.Title, task.Title, "Create() task = [%v], wantTask = [%v]", task, tt.wantTask)
			}
			s.TearDownTest()
		})
	}
}

func (s *bizTestSuite) Test_impl_List() {
	type args struct {
		page int32
		size int32
	}
	tests := []struct {
		name      string
		args      args
		mockFunc  func()
		wantTasks []*entities.Task
		wantErr   string
	}{
		{
			name: "list 1 3 then tasks[3] nil",
			args: args{1, 3},
			mockFunc: func() {
				s.taskRepo.On("QueryTaskList", mock.AnythingOfType("int32"), mock.AnythingOfType("int32")).Return(
					[]*entities.Task{
						{Title: "1"},
						{Title: "2"},
						{Title: "3"},
					}, nil).Once()
			},
			wantTasks: []*entities.Task{
				{Title: "1"},
				{Title: "2"},
				{Title: "3"},
			},
			wantErr: "",
		},
		{
			name: "list 1 3 then nil error",
			args: args{1, 3},
			mockFunc: func() {
				s.taskRepo.On("QueryTaskList", mock.AnythingOfType("int32"), mock.AnythingOfType("int32")).Return(
					nil, errors.New("test error")).Once()
			},
			wantTasks: nil,
			wantErr:   "test error",
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mockFunc()
			gotTasks, err := s.taskBiz.List(tt.args.page, tt.args.size)
			if err != nil {
				s.EqualErrorf(err, tt.wantErr, "List() error = %v, wantErr %v", err, tt.wantErr)
			}

			if gotTasks != nil {
				s.EqualValuesf(tt.wantTasks, gotTasks, "List() gotTasks = %v, want %v", gotTasks, tt.wantTasks)
			}
		})
	}
}

func (s *bizTestSuite) Test_impl_Remove() {
	type args struct {
		id string
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
		wantErr   error
		mockFunc  func()
	}{
		{
			name:      "empty then 0 error",
			args:      args{""},
			wantCount: 0,
			wantErr:   errors.New("id must be NOT empty"),
			mockFunc:  func() {},
		},
		{
			name:      "test then 0 error",
			args:      args{"test"},
			wantCount: 0,
			wantErr:   errors.New("invalid UUID length: 4"),
			mockFunc:  func() {},
		},
		{
			name:      "f3d58c97-e50e-4a00-ba51-ef7d2bec02e0 then 0 error",
			args:      args{"f3d58c97-e50e-4a00-ba51-ef7d2bec02e0"},
			wantCount: 0,
			wantErr:   errors.New("test error"),
			mockFunc: func() {
				s.taskRepo.On("RemoveTask", mock.AnythingOfType("string")).Return(
					0, errors.New("test error")).Once()
			},
		},
		{
			name:      "f3d58c97-e50e-4a00-ba51-ef7d2bec02e0 then 1 nil",
			args:      args{"f3d58c97-e50e-4a00-ba51-ef7d2bec02e0"},
			wantCount: 1,
			wantErr:   nil,
			mockFunc: func() {
				s.taskRepo.On("RemoveTask", mock.AnythingOfType("string")).Return(1, nil).Once()
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mockFunc()
			gotCount, err := s.taskBiz.Remove(tt.args.id)
			if err != nil || tt.wantErr != nil {
				s.EqualErrorf(err, tt.wantErr.Error(), "[%s] Remove() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
			s.EqualValuesf(tt.wantCount, gotCount, "[%s] Remove() gotCount = %v, want %v", tt.name, gotCount, tt.wantCount)

			s.TearDownTest()
		})
	}
}

func TestTaskBiz(t *testing.T) {
	suite.Run(t, new(bizTestSuite))
}
