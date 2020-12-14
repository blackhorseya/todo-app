package task

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blackhorseya/todo-app/internal/app/biz/task/mocks"
	"github.com/blackhorseya/todo-app/internal/app/entities"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type handlerTestSuite struct {
	suite.Suite
	r           *gin.Engine
	taskBiz     *mocks.Biz
	taskHandler IHandler
}

func (s *handlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	s.r = gin.New()

	s.taskBiz = new(mocks.Biz)
	handler, err := CreateTaskHandler(s.taskBiz)
	if err != nil {
		panic(err)
	}
	s.taskHandler = handler
}

func (s *handlerTestSuite) TearDownTest() {
	s.taskBiz.AssertExpectations(s.T())
}

func (s *handlerTestSuite) Test_impl_List() {
	s.r.GET("/api/v1/tasks", s.taskHandler.List)

	type args struct {
		page string
		size string
	}
	tests := []struct {
		name      string
		args      args
		mockFunc  func()
		wantCode  int
		wantTasks []*entities.Task
	}{
		{
			name: "list 10 10 then 204 nil",
			args: args{page: "10", size: "10"},
			mockFunc: func() {
				s.taskBiz.On("List", int32(10), int32(10)).Return(nil, nil).Once()
			},
			wantCode: http.StatusNoContent,
		},
		{
			name: "list 1 1 then 200 tasks",
			args: args{page: "1", size: "1"},
			mockFunc: func() {
				s.taskBiz.On("List", int32(1), int32(1)).Return([]*entities.Task{
					{Title: "test"},
				}, nil).Once()
			},
			wantCode: http.StatusOK,
			wantTasks: []*entities.Task{
				{Title: "test"},
			},
		},
		{
			name: "list a b then 400 nil",
			args: args{page: "a", size: "b"},
			mockFunc: func() {
				s.taskBiz.On("List", mock.AnythingOfType("int32"), mock.AnythingOfType("int32")).Return(
					nil, nil).Once()
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name: "list 10 b then 400 nil",
			args: args{page: "10", size: "b"},
			mockFunc: func() {
				s.taskBiz.On("List", mock.AnythingOfType("int32"), mock.AnythingOfType("int32")).Return(
					nil, nil).Once()
			},
			wantCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		tt.mockFunc()
		uri := fmt.Sprintf("/api/v1/tasks?page=%s&size=%s", tt.args.page, tt.args.size)
		req := httptest.NewRequest(http.MethodGet, uri, nil)
		w := httptest.NewRecorder()

		s.r.ServeHTTP(w, req)

		got := w.Result()
		defer func() {
			_ = got.Body.Close()
		}()

		body, _ := ioutil.ReadAll(got.Body)
		var gotTasks []*entities.Task
		err := json.Unmarshal(body, &gotTasks)
		if err != nil {
			s.Errorf(err, "unmarshal response body")
		}

		s.EqualValuesf(tt.wantCode, got.StatusCode, "List() code = [%v], wantCode = [%v]", got.StatusCode, tt.wantCode)
		if tt.wantTasks != nil {
			s.EqualValuesf(tt.wantTasks, gotTasks, "List() tasks = [%v], wantTasks = [%v]", gotTasks, tt.wantTasks)
		}
		s.TearDownTest()
	}
}

func (s *handlerTestSuite) Test_impl_Create() {
	s.r.POST("/api/v1/tasks", s.taskHandler.Create)

	type args struct {
		newTask *entities.Task
	}
	tests := []struct {
		name     string
		args     args
		mockFunc func()
		wantCode int
		wantTask *entities.Task
	}{
		{
			name: "create newTask then 201 task",
			args: args{&entities.Task{Title: "test"}},
			mockFunc: func() {
				s.taskBiz.On("Create", mock.AnythingOfType("*entities.Task")).Return(
					&entities.Task{Title: "test"}, nil).Once()
			},
			wantCode: http.StatusCreated,
			wantTask: &entities.Task{Title: "test"},
		},
		{
			name: "create nil then 400 nil",
			args: args{nil},
			mockFunc: func() {
				s.taskBiz.On("Create", mock.AnythingOfType("*entities.Task")).Return(
					nil, nil).Once()
			},
			wantCode: http.StatusBadRequest,
			wantTask: nil,
		},
		{
			name: "create newTask then 500 nil error",
			args: args{&entities.Task{Title: "test"}},
			mockFunc: func() {
				s.taskBiz.On("Create", mock.AnythingOfType("*entities.Task")).Return(
					nil, errors.New("test error")).Once()
			},
			wantCode: http.StatusInternalServerError,
			wantTask: nil,
		},
	}
	for _, tt := range tests {
		tt.mockFunc()
		uri := fmt.Sprintf("/api/v1/tasks")
		newTask, _ := json.Marshal(tt.args.newTask)
		req := httptest.NewRequest(http.MethodPost, uri, bytes.NewBuffer(newTask))
		w := httptest.NewRecorder()

		s.r.ServeHTTP(w, req)

		got := w.Result()
		defer func() {
			_ = got.Body.Close()
		}()

		body, _ := ioutil.ReadAll(got.Body)
		var gotTask *entities.Task
		err := json.Unmarshal(body, &gotTask)
		if err != nil {
			s.Errorf(err, "unmarshal response body")
		}

		s.EqualValuesf(tt.wantCode, got.StatusCode, "Create() code = [%v], wantCode = [%v]", got.StatusCode, tt.wantCode)
		if tt.wantTask != nil {
			s.EqualValuesf(tt.wantTask, gotTask, "Create() task = [%v], wantTask = [%v]", gotTask, tt.wantTask)
		}
		s.TearDownTest()
	}
}

func TestTaskHandler(t *testing.T) {
	suite.Run(t, new(handlerTestSuite))
}
