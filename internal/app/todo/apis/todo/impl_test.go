package todo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo/mocks"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/er"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"github.com/blackhorseya/todo-app/internal/pkg/infra/transports/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

var (
	uuid1 = int64(1)

	task1 = &todo.Task{
		Id:    uuid1,
		Title: "title",
	}

	updated1 = &todo.Task{
		Id:        uuid1,
		Completed: true,
	}

	updated2 = &todo.Task{
		Id:    uuid1,
		Title: "title",
	}
)

type handlerSuite struct {
	suite.Suite
	r       *gin.Engine
	mock    *mocks.IBiz
	handler IHandler
}

func (s *handlerSuite) SetupTest() {
	logger, _ := zap.NewDevelopment()

	gin.SetMode(gin.TestMode)
	s.r = gin.New()
	s.r.Use(middlewares.ContextMiddleware())
	s.r.Use(middlewares.ErrorMiddleware())

	s.mock = new(mocks.IBiz)
	handler, err := CreateIHandler(logger, s.mock)
	if err != nil {
		panic(err)
	}

	s.handler = handler
}

func (s *handlerSuite) TearDownTest() {
	s.mock.AssertExpectations(s.T())
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(handlerSuite))
}

func (s *handlerSuite) Test_impl_GetByID() {
	s.r.GET("/api/v1/tasks/:id", s.handler.GetByID)

	type args struct {
		id   int64
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "get by id then error",
			args: args{id: uuid1, mock: func() {
				s.mock.On("GetByID", mock.Anything, uuid1).Return(nil, er.ErrGetTask).Once()
			}},
			wantCode: 500,
		},
		{
			name: "get by id then not found error",
			args: args{id: uuid1, mock: func() {
				s.mock.On("GetByID", mock.Anything, uuid1).Return(nil, er.ErrTaskNotExists).Once()
			}},
			wantCode: 404,
		},
		{
			name: "get by id then success",
			args: args{id: uuid1, mock: func() {
				s.mock.On("GetByID", mock.Anything, uuid1).Return(task1, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/tasks/%v", tt.args.id)
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "GetByID() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_List() {
	s.r.GET("/api/v1/tasks", s.handler.List)

	type args struct {
		start string
		end   string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "start not a integer then error",
			args:     args{start: "start", end: "3"},
			wantCode: 400,
		},
		{
			name:     "end not a integer then error",
			args:     args{start: "0", end: "end"},
			wantCode: 400,
		},
		{
			name: "list then error",
			args: args{start: "0", end: "3", mock: func() {
				s.mock.On("List", mock.Anything, 0, 3).Return(nil, 0, er.ErrListTasks).Once()
			}},
			wantCode: 500,
		},
		{
			name: "list then not found error",
			args: args{start: "0", end: "3", mock: func() {
				s.mock.On("List", mock.Anything, 0, 3).Return(nil, 0, er.ErrTaskNotExists).Once()
			}},
			wantCode: 404,
		},
		{
			name: "list then success",
			args: args{start: "0", end: "3", mock: func() {
				s.mock.On("List", mock.Anything, 0, 3).Return([]*todo.Task{task1}, 10, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/tasks?start=%v&end=%v", tt.args.start, tt.args.end)
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "List() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_Create() {
	s.r.POST("/api/v1/tasks", s.handler.Create)

	type args struct {
		created *todo.Task
		mock    func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "create then error",
			args: args{created: task1, mock: func() {
				s.mock.On("Create", mock.Anything, task1.Title).Return(nil, er.ErrCreateTask).Once()
			}},
			wantCode: 500,
		},
		{
			name: "create then success",
			args: args{created: task1, mock: func() {
				s.mock.On("Create", mock.Anything, task1.Title).Return(task1, nil).Once()
			}},
			wantCode: 201,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/tasks")
			data, _ := json.Marshal(tt.args.created)
			req := httptest.NewRequest(http.MethodPost, uri, bytes.NewBuffer(data))
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Create() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_UpdateStatus() {
	s.r.PATCH("/api/v1/tasks/:id/status", s.handler.UpdateStatus)

	type args struct {
		id      int64
		updated *todo.Task
		mock    func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "update status then error",
			args: args{id: uuid1, updated: updated1, mock: func() {
				s.mock.On("UpdateStatus", mock.Anything, uuid1, updated1.Completed).Return(nil, er.ErrUpdateStatusTask).Once()
			}},
			wantCode: 500,
		},
		{
			name: "update status then success",
			args: args{id: uuid1, updated: updated1, mock: func() {
				s.mock.On("UpdateStatus", mock.Anything, uuid1, updated1.Completed).Return(updated1, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/tasks/%v/status", tt.args.id)
			data, _ := json.Marshal(tt.args.updated)
			req := httptest.NewRequest(http.MethodPatch, uri, bytes.NewBuffer(data))
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "UpdateStatus() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_ChangeTitle() {
	s.r.PATCH("/api/v1/tasks/:id/title", s.handler.ChangeTitle)

	type args struct {
		id      int64
		updated *todo.Task
		mock    func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "change title then error",
			args: args{id: uuid1, updated: updated1, mock: func() {
				s.mock.On("ChangeTitle", mock.Anything, uuid1, updated1.Title).Return(nil, er.ErrChangeTitleTask).Once()
			}},
			wantCode: 500,
		},
		{
			name: "change title then success",
			args: args{id: uuid1, updated: updated1, mock: func() {
				s.mock.On("ChangeTitle", mock.Anything, uuid1, updated1.Title).Return(updated1, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/tasks/%v/title", tt.args.id)
			data, _ := json.Marshal(tt.args.updated)
			req := httptest.NewRequest(http.MethodPatch, uri, bytes.NewBuffer(data))
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "ChangeTitle() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_Delete() {
	s.r.DELETE("/api/v1/tasks/:id", s.handler.Delete)

	type args struct {
		id   int64
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "delete then error",
			args: args{id: uuid1, mock: func() {
				s.mock.On("Delete", mock.Anything, uuid1).Return(er.ErrDeleteTask).Once()
			}},
			wantCode: 500,
		},
		{
			name: "delete then success",
			args: args{id: uuid1, mock: func() {
				s.mock.On("Delete", mock.Anything, uuid1).Return(nil).Once()
			}},
			wantCode: 204,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/tasks/%v", tt.args.id)
			req := httptest.NewRequest(http.MethodDelete, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Delete() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}
