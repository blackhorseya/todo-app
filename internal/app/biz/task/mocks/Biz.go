// Code generated by mockery v2.4.0. DO NOT EDIT.

package mocks

import (
	entities "github.com/blackhorseya/todo-app/internal/app/entities"
	mock "github.com/stretchr/testify/mock"
)

// Biz is an autogenerated mock type for the Biz type
type Biz struct {
	mock.Mock
}

// ChangeTitle provides a mock function with given fields: id, newTitle
func (_m *Biz) ChangeTitle(id string, newTitle string) (*entities.Task, error) {
	ret := _m.Called(id, newTitle)

	var r0 *entities.Task
	if rf, ok := ret.Get(0).(func(string, string) *entities.Task); ok {
		r0 = rf(id, newTitle)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, newTitle)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: newTask
func (_m *Biz) Create(newTask *entities.Task) (*entities.Task, error) {
	ret := _m.Called(newTask)

	var r0 *entities.Task
	if rf, ok := ret.Get(0).(func(*entities.Task) *entities.Task); ok {
		r0 = rf(newTask)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entities.Task) error); ok {
		r1 = rf(newTask)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: page, size
func (_m *Biz) List(page int32, size int32) ([]*entities.Task, error) {
	ret := _m.Called(page, size)

	var r0 []*entities.Task
	if rf, ok := ret.Get(0).(func(int32, int32) []*entities.Task); ok {
		r0 = rf(page, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32, int32) error); ok {
		r1 = rf(page, size)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: id
func (_m *Biz) Remove(id string) (bool, error) {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: id, completed
func (_m *Biz) UpdateStatus(id string, completed bool) (*entities.Task, error) {
	ret := _m.Called(id, completed)

	var r0 *entities.Task
	if rf, ok := ret.Get(0).(func(string, bool) *entities.Task); ok {
		r0 = rf(id, completed)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(id, completed)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}