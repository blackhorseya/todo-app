// Code generated by mockery v2.7.4. DO NOT EDIT.

package mocks

import (
	contextx "github.com/blackhorseya/todo-app/internal/pkg/base/contextx"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// IRepo is an autogenerated mock type for the IRepo type
type IRepo struct {
	mock.Mock
}

// Ping provides a mock function with given fields: ctx, timeout
func (_m *IRepo) Ping(ctx contextx.Contextx, timeout time.Duration) error {
	ret := _m.Called(ctx, timeout)

	var r0 error
	if rf, ok := ret.Get(0).(func(contextx.Contextx, time.Duration) error); ok {
		r0 = rf(ctx, timeout)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
