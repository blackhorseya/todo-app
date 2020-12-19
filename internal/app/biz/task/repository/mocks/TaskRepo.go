// Code generated by mockery v2.4.0. DO NOT EDIT.

package mocks

import (
	entities "github.com/blackhorseya/todo-app/internal/app/entities"
	mock "github.com/stretchr/testify/mock"
)

// TaskRepo is an autogenerated mock type for the TaskRepo type
type TaskRepo struct {
	mock.Mock
}

// CreateTask provides a mock function with given fields: newTask
func (_m *TaskRepo) CreateTask(newTask *entities.Task) (*entities.Task, error) {
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

// FindOne provides a mock function with given fields: id
func (_m *TaskRepo) FindOne(id string) (*entities.Task, error) {
	ret := _m.Called(id)

	var r0 *entities.Task
	if rf, ok := ret.Get(0).(func(string) *entities.Task); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryTaskList provides a mock function with given fields: limit, offset
func (_m *TaskRepo) QueryTaskList(limit int32, offset int32) ([]*entities.Task, error) {
	ret := _m.Called(limit, offset)

	var r0 []*entities.Task
	if rf, ok := ret.Get(0).(func(int32, int32) []*entities.Task); ok {
		r0 = rf(limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32, int32) error); ok {
		r1 = rf(limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveTask provides a mock function with given fields: id
func (_m *TaskRepo) RemoveTask(id string) (int, error) {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTask provides a mock function with given fields: updated
func (_m *TaskRepo) UpdateTask(updated *entities.Task) (*entities.Task, error) {
	ret := _m.Called(updated)

	var r0 *entities.Task
	if rf, ok := ret.Get(0).(func(*entities.Task) *entities.Task); ok {
		r0 = rf(updated)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entities.Task) error); ok {
		r1 = rf(updated)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
