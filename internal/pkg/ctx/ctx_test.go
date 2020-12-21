package ctx

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ctxSuite struct {
	suite.Suite
}

func TestCTXSuite(t *testing.T) {
	suite.Run(t, new(ctxSuite))
}

func (s *ctxSuite) TestBackground() {
	tests := []struct {
		name string
		want CTX
	}{
		{
			name: "background then background",
			want: Background(),
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			if got := Background(); !reflect.DeepEqual(got, tt.want) {
				s.T().Errorf("Background() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *ctxSuite) TestWithCancel() {
	ctx, cancel := WithCancel(Background())
	defer cancel()

	type args struct {
		parent CTX
	}
	tests := []struct {
		name  string
		args  args
		want  CTX
		want1 context.CancelFunc
	}{
		{
			name:  "withCancel then withCancel",
			args:  args{Background()},
			want:  ctx,
			want1: cancel,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			got, _ := WithCancel(tt.args.parent)
			if !reflect.DeepEqual(got, tt.want) {
				s.T().Errorf("WithCancel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *ctxSuite) TestWithTimeout() {
	ctx, cancel := WithTimeout(Background(), 10*time.Second)
	defer cancel()

	type args struct {
		parent CTX
		d      time.Duration
	}
	tests := []struct {
		name  string
		args  args
		want  CTX
		want1 context.CancelFunc
	}{
		{
			name:  "withTimeout then withTimeout",
			args:  args{Background(), 10 * time.Second},
			want:  ctx,
			want1: cancel,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			_, _ = WithTimeout(tt.args.parent, tt.args.d)
		})
	}
}

func (s *ctxSuite) TestWithValue() {
	ctx := WithValue(Background(), "key", "value")

	type args struct {
		parent CTX
		key    string
		val    interface{}
	}
	tests := []struct {
		name string
		args args
		want CTX
	}{
		{
			name: "withValue then withValue",
			args: args{Background(), "key", "value"},
			want: ctx,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			if got := WithValue(tt.args.parent, tt.args.key, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				s.T().Errorf("WithValue() = %v, want %v", got, tt.want)
			}
		})
	}
}