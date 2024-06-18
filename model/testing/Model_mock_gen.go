// Code generated by mockery v2.40.3. DO NOT EDIT.

package modeltesting

import (
	mock "github.com/stretchr/testify/mock"
	metrics "github.com/symflower/eval-dev-quality/evaluate/metrics"

	model "github.com/symflower/eval-dev-quality/model"

	task "github.com/symflower/eval-dev-quality/task"
)

// MockModel is an autogenerated mock type for the Model type
type MockModel struct {
	mock.Mock
}

// ID provides a mock function with given fields:
func (_m *MockModel) ID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsTaskSupported provides a mock function with given fields: taskIdentifier
func (_m *MockModel) IsTaskSupported(taskIdentifier task.Identifier) bool {
	ret := _m.Called(taskIdentifier)

	if len(ret) == 0 {
		panic("no return value specified for IsTaskSupported")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(task.Identifier) bool); ok {
		r0 = rf(taskIdentifier)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RunTask provides a mock function with given fields: ctx, taskIdentifier
func (_m *MockModel) RunTask(ctx model.Context, taskIdentifier task.Identifier) (metrics.Assessments, error) {
	ret := _m.Called(ctx, taskIdentifier)

	if len(ret) == 0 {
		panic("no return value specified for RunTask")
	}

	var r0 metrics.Assessments
	var r1 error
	if rf, ok := ret.Get(0).(func(model.Context, task.Identifier) (metrics.Assessments, error)); ok {
		return rf(ctx, taskIdentifier)
	}
	if rf, ok := ret.Get(0).(func(model.Context, task.Identifier) metrics.Assessments); ok {
		r0 = rf(ctx, taskIdentifier)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metrics.Assessments)
		}
	}

	if rf, ok := ret.Get(1).(func(model.Context, task.Identifier) error); ok {
		r1 = rf(ctx, taskIdentifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockModel creates a new instance of MockModel. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockModel(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockModel {
	mock := &MockModel{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
