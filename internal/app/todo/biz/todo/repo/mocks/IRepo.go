// Code generated by mockery v2.7.4. DO NOT EDIT.

package mocks

import (
	contextx "github.com/blackhorseya/todo-app/internal/pkg/base/contextx"
	mock "github.com/stretchr/testify/mock"

	todo "github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
)

// IRepo is an autogenerated mock type for the IRepo type
type IRepo struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx
func (_m *IRepo) Count(ctx contextx.Contextx) (int, error) {
	ret := _m.Called(ctx)

	var r0 int
	if rf, ok := ret.Get(0).(func(contextx.Contextx) int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, title
func (_m *IRepo) Create(ctx contextx.Contextx, title string) (*todo.Task, error) {
	ret := _m.Called(ctx, title)

	var r0 *todo.Task
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string) *todo.Task); ok {
		r0 = rf(ctx, title)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*todo.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string) error); ok {
		r1 = rf(ctx, title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *IRepo) GetByID(ctx contextx.Contextx, id string) (*todo.Task, error) {
	ret := _m.Called(ctx, id)

	var r0 *todo.Task
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string) *todo.Task); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*todo.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, limit, offset
func (_m *IRepo) List(ctx contextx.Contextx, limit int, offset int) ([]*todo.Task, error) {
	ret := _m.Called(ctx, limit, offset)

	var r0 []*todo.Task
	if rf, ok := ret.Get(0).(func(contextx.Contextx, int, int) []*todo.Task); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*todo.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, int, int) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: ctx, id
func (_m *IRepo) Remove(ctx contextx.Contextx, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, updated
func (_m *IRepo) Update(ctx contextx.Contextx, updated *todo.Task) (*todo.Task, error) {
	ret := _m.Called(ctx, updated)

	var r0 *todo.Task
	if rf, ok := ret.Get(0).(func(contextx.Contextx, *todo.Task) *todo.Task); ok {
		r0 = rf(ctx, updated)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*todo.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, *todo.Task) error); ok {
		r1 = rf(ctx, updated)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
