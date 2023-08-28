// Code generated by mockery v2.33.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/modaniru/avito/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// History is an autogenerated mock type for the History type
type History struct {
	mock.Mock
}

// GetHistory provides a mock function with given fields: ctx
func (_m *History) GetHistory(ctx context.Context) ([]entity.History, error) {
	ret := _m.Called(ctx)

	var r0 []entity.History
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]entity.History, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entity.History); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.History)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHistoryByDate provides a mock function with given fields: ctx, date
func (_m *History) GetHistoryByDate(ctx context.Context, date string) ([]entity.History, error) {
	ret := _m.Called(ctx, date)

	var r0 []entity.History
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]entity.History, error)); ok {
		return rf(ctx, date)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []entity.History); ok {
		r0 = rf(ctx, date)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.History)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, date)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewHistory creates a new instance of History. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHistory(t interface {
	mock.TestingT
	Cleanup(func())
}) *History {
	mock := &History{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
