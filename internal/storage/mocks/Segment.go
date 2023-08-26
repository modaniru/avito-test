// Code generated by mockery v2.33.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/modaniru/avito/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// Segment is an autogenerated mock type for the Segment type
type Segment struct {
	mock.Mock
}

// DeleteSegment provides a mock function with given fields: ctx, name
func (_m *Segment) DeleteSegment(ctx context.Context, name string) error {
	ret := _m.Called(ctx, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSegments provides a mock function with given fields: ctx
func (_m *Segment) GetSegments(ctx context.Context) ([]entity.Segment, error) {
	ret := _m.Called(ctx)

	var r0 []entity.Segment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]entity.Segment, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Segment); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Segment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveSegment provides a mock function with given fields: ctx, name
func (_m *Segment) SaveSegment(ctx context.Context, name string) (int, error) {
	ret := _m.Called(ctx, name)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSegment creates a new instance of Segment. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSegment(t interface {
	mock.TestingT
	Cleanup(func())
}) *Segment {
	mock := &Segment{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
