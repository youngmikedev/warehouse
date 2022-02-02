// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/imranzahaev/warehouse/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// User is an autogenerated mock type for the User type
type User struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, user, password
func (_m *User) Create(ctx context.Context, user domain.User, password string) (int, error) {
	ret := _m.Called(ctx, user, password)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, domain.User, string) int); ok {
		r0 = rf(ctx, user, password)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.User, string) error); ok {
		r1 = rf(ctx, user, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSession provides a mock function with given fields: ctx, session
func (_m *User) CreateSession(ctx context.Context, session domain.Session) (int, error) {
	ret := _m.Called(ctx, session)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, domain.Session) int); ok {
		r0 = rf(ctx, session)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Session) error); ok {
		r1 = rf(ctx, session)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, id
func (_m *User) Get(ctx context.Context, id int) (domain.User, string, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(context.Context, int) domain.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, int) string); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, int) error); ok {
		r2 = rf(ctx, id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByLogin provides a mock function with given fields: ctx, login
func (_m *User) GetByLogin(ctx context.Context, login string) (domain.User, string, error) {
	ret := _m.Called(ctx, login)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.User); ok {
		r0 = rf(ctx, login)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, string) string); ok {
		r1 = rf(ctx, login)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, login)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetSessionByAccess provides a mock function with given fields: ctx, token
func (_m *User) GetSessionByAccess(ctx context.Context, token string) (domain.Session, error) {
	ret := _m.Called(ctx, token)

	var r0 domain.Session
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Session); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(domain.Session)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSessionByRefresh provides a mock function with given fields: ctx, token
func (_m *User) GetSessionByRefresh(ctx context.Context, token string) (domain.Session, error) {
	ret := _m.Called(ctx, token)

	var r0 domain.Session
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Session); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(domain.Session)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, user, password
func (_m *User) Update(ctx context.Context, user domain.User, password string) error {
	ret := _m.Called(ctx, user, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.User, string) error); ok {
		r0 = rf(ctx, user, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateSession provides a mock function with given fields: ctx, session
func (_m *User) UpdateSession(ctx context.Context, session domain.Session) error {
	ret := _m.Called(ctx, session)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Session) error); ok {
		r0 = rf(ctx, session)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
