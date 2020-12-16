package health

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blackhorseya/todo-app/internal/app/biz/health/mocks"
	"github.com/blackhorseya/todo-app/internal/app/entities"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type healthTestSuite struct {
	suite.Suite
	r       *gin.Engine
	biz     *mocks.Biz
	handler IHandler
}

func (s *healthTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	s.r = gin.New()
	s.biz = new(mocks.Biz)
	handler, err := CreateHealthHandler(s.biz)
	if err != nil {
		panic(err)
	}

	s.handler = handler
}

func (s *healthTestSuite) TearDownTest() {
	s.biz.AssertExpectations(s.T())
}

func (s *healthTestSuite) Test_impl_Readiness() {
	s.r.GET("/api/readiness", s.handler.Readiness)

	type args struct {
	}
	tests := []struct {
		name     string
		args     args
		mockFunc func()
		wantCode int
		wantRes  *entities.Response
	}{
		{
			name: "readiness then 200 ok",
			args: args{},
			mockFunc: func() {
				s.biz.On("Readiness").Return(true, nil).Once()
			},
			wantCode: http.StatusOK,
			wantRes: &entities.Response{
				Ok:  true,
				Msg: "application has been ready",
			},
		},
		{
			name: "readiness then 500 fail",
			args: args{},
			mockFunc: func() {
				s.biz.On("Readiness").Return(
					false, errors.New("test error")).Once()
			},
			wantCode: http.StatusInternalServerError,
			wantRes: &entities.Response{
				Ok:  false,
				Msg: "application has failed",
			},
		},
	}
	for _, tt := range tests {
		tt.mockFunc()
		uri := "/api/readiness"
		req := httptest.NewRequest(http.MethodGet, uri, nil)
		w := httptest.NewRecorder()

		s.r.ServeHTTP(w, req)

		got := w.Result()
		defer got.Body.Close()

		body, _ := ioutil.ReadAll(got.Body)
		var gotRes *entities.Response
		err := json.Unmarshal(body, &gotRes)
		if err != nil {
			s.Errorf(err, "unmarshal response body")
		}

		s.EqualValuesf(tt.wantCode, got.StatusCode, "Readiness() code = [%v], wantCode = [%v]", got.StatusCode, tt.wantCode)
		s.EqualValuesf(tt.wantRes, gotRes, "Readiness() res = [%v], wantRes = [%v]", gotRes, tt.wantRes)

		s.TearDownTest()
	}
}

func (s *healthTestSuite) Test_impl_Liveness() {
	s.r.GET("/api/liveness", s.handler.Liveness)

	type args struct {
	}
	tests := []struct {
		name     string
		args     args
		mockFunc func()
		wantCode int
		wantRes  *entities.Response
	}{
		{
			name: "liveness then 200 ok",
			args: args{},
			mockFunc: func() {
				s.biz.On("Liveness").Return(true, nil).Once()
			},
			wantCode: http.StatusOK,
			wantRes: &entities.Response{
				Ok:  true,
				Msg: "alive",
			},
		},
		{
			name: "liveness then 500 fail",
			args: args{},
			mockFunc: func() {
				s.biz.On("Liveness").Return(
					false, errors.New("test error")).Once()
			},
			wantCode: http.StatusInternalServerError,
			wantRes: &entities.Response{
				Ok:  false,
				Msg: "dead",
			},
		},
	}
	for _, tt := range tests {
		tt.mockFunc()
		uri := "/api/liveness"
		req := httptest.NewRequest(http.MethodGet, uri, nil)
		w := httptest.NewRecorder()

		s.r.ServeHTTP(w, req)

		got := w.Result()
		defer got.Body.Close()

		body, _ := ioutil.ReadAll(got.Body)
		var gotRes *entities.Response
		err := json.Unmarshal(body, &gotRes)
		if err != nil {
			s.Errorf(err, "unmarshal response body")
		}

		s.EqualValuesf(tt.wantCode, got.StatusCode, "Liveness() code = [%v], wantCode = [%v]", got.StatusCode, tt.wantCode)
		s.EqualValuesf(tt.wantRes, gotRes, "Liveness() res = [%v], wantRes = [%v]", gotRes, tt.wantRes)

		s.TearDownTest()
	}
}

func TestHealthHandler(t *testing.T) {
	suite.Run(t, new(healthTestSuite))
}
