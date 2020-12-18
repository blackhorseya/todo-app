package health

import (
	"errors"
	"testing"

	"github.com/blackhorseya/todo-app/internal/app/biz/health/repository/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type bizTestSuite struct {
	suite.Suite
	healthRepo *mocks.HealthRepo
	healthBiz  Biz
}

func (s *bizTestSuite) SetupTest() {
	s.healthRepo = new(mocks.HealthRepo)
	biz, err := CreateHealthBiz(s.healthRepo)
	if err != nil {
		panic(err)
	}

	s.healthBiz = biz
}

func (s *bizTestSuite) TearDownTest() {
	s.healthRepo.AssertExpectations(s.T())
}

func (s *bizTestSuite) Test_impl_Readiness() {
	tests := []struct {
		name     string
		wantOk   bool
		wantErr  string
		mockFunc func()
	}{
		{
			name:    "no error then true nil",
			wantOk:  true,
			wantErr: "",
			mockFunc: func() {
				s.healthRepo.On("Ping", mock.AnythingOfType("time.Duration")).Return(nil).Once()
			},
		},
		{
			name:    "has error then false error",
			wantOk:  false,
			wantErr: "test error",
			mockFunc: func() {
				s.healthRepo.On("Ping", mock.AnythingOfType("time.Duration")).Return(
					errors.New("test error")).Once()
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mockFunc()
			gotOk, err := s.healthBiz.Readiness()
			if err != nil {
				s.EqualErrorf(err, tt.wantErr, "Readiness() error = %v, wantErr %v", err, tt.wantErr)
			}
			s.EqualValuesf(tt.wantOk, gotOk, "Readiness() gotOk = %v, want %v", gotOk, tt.wantOk)
			s.TearDownTest()
		})
	}
}

func (s *bizTestSuite) Test_impl_Liveness() {
	tests := []struct {
		name     string
		wantOk   bool
		wantErr  string
		mockFunc func()
	}{
		{
			name:    "no error then true nil",
			wantOk:  true,
			wantErr: "",
			mockFunc: func() {
				s.healthRepo.On("Ping", mock.AnythingOfType("time.Duration")).Return(nil).Once()
			},
		},
		{
			name:    "has error then false error",
			wantOk:  false,
			wantErr: "test error",
			mockFunc: func() {
				s.healthRepo.On("Ping", mock.AnythingOfType("time.Duration")).Return(
					errors.New("test error")).Once()
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mockFunc()
			gotOk, err := s.healthBiz.Liveness()
			if err != nil {
				s.EqualErrorf(err, tt.wantErr, "Liveness() error = %v, wantErr %v", err, tt.wantErr)
			}
			s.EqualValuesf(tt.wantOk, gotOk, "Liveness() gotOk = %v, want %v", gotOk, tt.wantOk)
			s.TearDownTest()
		})
	}
}

func TestHealthBiz(t *testing.T) {
	suite.Run(t, new(bizTestSuite))
}
