// Code generated by mockery v2.33.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/modaniru/avito/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// User is an autogenerated mock type for the User type
type User struct {
	mock.Mock
}

// DeleteUser provides a mock function with given fields: ctx, userId
func (_m *User) DeleteUser(ctx context.Context, userId int) error {
	ret := _m.Called(ctx, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FollowRandomUsers provides a mock function with given fields: ctx, name, percent
func (_m *User) FollowRandomUsers(ctx context.Context, name string, percent float64) (int, error) {
	ret := _m.Called(ctx, name, percent)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, float64) (int, error)); ok {
		return rf(ctx, name, percent)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, float64) int); ok {
		r0 = rf(ctx, name, percent)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, float64) error); ok {
		r1 = rf(ctx, name, percent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FollowToSegments provides a mock function with given fields: ctx, userId, segments, date
func (_m *User) FollowToSegments(ctx context.Context, userId int, segments []string, date *string) error {
	ret := _m.Called(ctx, userId, segments, date)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, []string, *string) error); ok {
		r0 = rf(ctx, userId, segments, date)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserSegments provides a mock function with given fields: ctx, id
func (_m *User) GetUserSegments(ctx context.Context, id int) ([]entity.Follows, error) {
	ret := _m.Called(ctx, id)

	var r0 []entity.Follows
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]entity.Follows, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []entity.Follows); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Follows)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields: ctx
func (_m *User) GetUsers(ctx context.Context) ([]entity.User, error) {
	ret := _m.Called(ctx)

	var r0 []entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]entity.User, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entity.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveUser provides a mock function with given fields: ctx, userId
func (_m *User) SaveUser(ctx context.Context, userId int) error {
	ret := _m.Called(ctx, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnFollowToSegments provides a mock function with given fields: ctx, userId, segments
func (_m *User) UnFollowToSegments(ctx context.Context, userId int, segments []string) error {
	ret := _m.Called(ctx, userId, segments)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, []string) error); ok {
		r0 = rf(ctx, userId, segments)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUser creates a new instance of User. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUser(t interface {
	mock.TestingT
	Cleanup(func())
}) *User {
	mock := &User{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
