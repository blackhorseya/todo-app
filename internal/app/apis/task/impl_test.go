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
	"google.golang.org/protobuf/types/known/anypb"
)

type taskTestSuite struct {
	suite.Suite
	r           *gin.Engine
	mockBiz     *mocks.Biz
	taskHandler IHandler
}

func (s *taskTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	s.r = gin.New()

	s.mockBiz = new(mocks.Biz)
	handler, err := CreateTaskHandler(s.mockBiz)
	if err != nil {
		panic(err)
	}
	s.taskHandler = handler
}

func (s *taskTestSuite) TearDownTest() {
	s.mockBiz.AssertExpectations(s.T())
}

func (s *taskTestSuite) Test_impl_List() {
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
			name: "list 10 10 then 404 nil",
			args: args{page: "10", size: "10"},
			mockFunc: func() {
				s.mockBiz.On("List", int32(10), int32(10)).Return(nil, nil).Once()
			},
			wantCode: http.StatusNotFound,
		},
		{
			name: "list 1 1 then 200 tasks",
			args: args{page: "1", size: "1"},
			mockFunc: func() {
				s.mockBiz.On("List", int32(1), int32(1)).Return([]*entities.Task{
					{Title: "test"},
				}, nil).Once()
				s.mockBiz.On("Count").Return(1, nil).Once()
			},
			wantCode: http.StatusOK,
			wantTasks: []*entities.Task{
				{Title: "test"},
			},
		},
		{
			name:     "list a b then 400 nil",
			args:     args{page: "a", size: "b"},
			mockFunc: func() {},
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "list 10 b then 400 nil",
			args:     args{page: "10", size: "b"},
			mockFunc: func() {},
			wantCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
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
				// s.EqualValuesf(tt.wantTasks, gotTasks, "List() tasks = [%v], wantTasks = [%v]", gotTasks, tt.wantTasks)
			}
			s.TearDownTest()
		})
	}
}

func (s *taskTestSuite) Test_impl_Create() {
	res1, _ := anypb.New(&entities.Task{Title: "test"})

	s.r.POST("/api/v1/tasks", s.taskHandler.Create)

	type args struct {
		newTask *entities.Task
	}
	tests := []struct {
		name     string
		args     args
		mockFunc func()
		wantCode int
		wantRes  *entities.Response
	}{
		{
			name: "create newTask then 201 task",
			args: args{&entities.Task{Title: "test"}},
			mockFunc: func() {
				s.mockBiz.On("Create", mock.AnythingOfType("*entities.Task")).Return(
					&entities.Task{Title: "test"}, nil).Once()
			},
			wantCode: http.StatusCreated,
			wantRes: &entities.Response{
				Ok:   true,
				Data: res1,
			},
		},
		{
			name:     "create nil then 400 nil",
			args:     args{nil},
			mockFunc: func() {},
			wantCode: http.StatusBadRequest,
			wantRes: &entities.Response{
				Ok:  false,
				Msg: "missing new task",
			},
		},
		{
			name: "create newTask then 500 nil error",
			args: args{&entities.Task{Title: "500"}},
			mockFunc: func() {
				s.mockBiz.On("Create", mock.AnythingOfType("*entities.Task")).Return(
					nil, errors.New("")).Once()
			},
			wantCode: http.StatusInternalServerError,
			wantRes: &entities.Response{
				Ok:  false,
				Msg: "",
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
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
			var gotRes *entities.Response
			err := json.Unmarshal(body, &gotRes)
			if err != nil {
				s.Errorf(err, "unmarshal response body")
			}

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Create() code = [%v], wantCode = [%v]", got.StatusCode, tt.wantCode)
			// s.EqualValuesf(tt.wantRes, gotRes, "Create() res = [%v], wantRes = [%v]", gotRes, tt.wantRes)

			s.TearDownTest()
		})
	}
}

func (s *taskTestSuite) Test_impl_Remove() {
	s.r.DELETE("/api/v1/tasks/:id", s.taskHandler.Remove)

	type args struct {
		id string
	}
	tests := []struct {
		name     string
		args     args
		mockFunc func()
		wantCode int
	}{
		{
			name:     "missing id then 404 error",
			args:     args{},
			mockFunc: func() {},
			wantCode: 404,
		},
		{
			name:     "123 then 400 error",
			args:     args{"123"},
			mockFunc: func() {},
			wantCode: 400,
		},
		{
			name: "f3d58c97-e50e-4a00-ba51-ef7d2bec02e0 then 200 nil",
			args: args{"f3d58c97-e50e-4a00-ba51-ef7d2bec02e0"},
			mockFunc: func() {
				s.mockBiz.On("Remove", "f3d58c97-e50e-4a00-ba51-ef7d2bec02e0").Return(
					1, nil).Once()
			},
			wantCode: 200,
		},
		{
			name: "f3d58c97-e50e-4a00-ba51-ef7d2bec02e0 then 500 error",
			args: args{"f3d58c97-e50e-4a00-ba51-ef7d2bec02e0"},
			mockFunc: func() {
				s.mockBiz.On("Remove", "f3d58c97-e50e-4a00-ba51-ef7d2bec02e0").Return(
					0, errors.New("test error"))
			},
			wantCode: 500,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mockFunc()
			uri := fmt.Sprintf("/api/v1/tasks/%s", tt.args.id)
			req := httptest.NewRequest(http.MethodDelete, uri, nil)
			w := httptest.NewRecorder()

			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Remove() code = [%v], wantCode = [%v]", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *taskTestSuite) Test_impl_ModifyInfo() {
	s.r.PATCH("/api/v1/tasks/:id", s.taskHandler.ModifyInfo)

	type args struct {
		id        string
		completed int
		title     string
		body      *modifyBody
	}
	tests := []struct {
		name     string
		args     args
		mockFunc func()
		wantCode int
	}{
		{
			name: "missing id then 404",
			args: args{id: ""},
			mockFunc: func() {

			},
			wantCode: 404,
		},
		{
			name: "id then 404",
			args: args{id: "id"},
			mockFunc: func() {

			},
			wantCode: 404,
		},
		{
			name: "uuid then 400",
			args: args{id: "f3d58c97-e50e-4a00-ba51-ef7d2bec02e0"},
			mockFunc: func() {

			},
			wantCode: 400,
		},
		{
			name: "uuid completed then 200",
			args: args{id: "f3d58c97-e50e-4a00-ba51-ef7d2bec02e0", completed: 1},
			mockFunc: func() {
				s.mockBiz.On("UpdateStatus", "f3d58c97-e50e-4a00-ba51-ef7d2bec02e0", false).Return(
					&entities.Task{
						Id:        "f3d58c97-e50e-4a00-ba51-ef7d2bec02e0",
						Completed: false,
					}, nil).Once()
			},
			wantCode: 200,
		},
		{
			name: "uuid title then 200",
			args: args{id: "f3d58c97-e50e-4a00-ba51-ef7d2bec02e0", title: "title"},
			mockFunc: func() {
				s.mockBiz.On("ChangeTitle", "f3d58c97-e50e-4a00-ba51-ef7d2bec02e0", "title").Return(
					&entities.Task{
						Id:    "f3d58c97-e50e-4a00-ba51-ef7d2bec02e0",
						Title: "title",
					}, nil).Once()
			},
			wantCode: 200,
		},
		{
			name: "uuid title empty then 400",
			args: args{id: "f3d58c97-e50e-4a00-ba51-ef7d2bec02e0", title: ""},
			mockFunc: func() {

			},
			wantCode: 400,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mockFunc()

			uri := fmt.Sprintf("/api/v1/tasks/%s?completed=%v&title=%v", tt.args.id, tt.args.completed, tt.args.title)
			req := httptest.NewRequest(http.MethodPatch, uri, nil)
			w := httptest.NewRecorder()

			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "ModifyInfo() code = [%v], wantCode = [%v]", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func TestTaskHandler(t *testing.T) {
	suite.Run(t, new(taskTestSuite))
}
