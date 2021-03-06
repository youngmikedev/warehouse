// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// TokenManager is an autogenerated mock type for the TokenManager type
type TokenManager struct {
	mock.Mock
}

// NewAccessToken provides a mock function with given fields:
func (_m *TokenManager) NewAccessToken() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewRefreshToken provides a mock function with given fields:
func (_m *TokenManager) NewRefreshToken() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ValidateAccessToken provides a mock function with given fields: createdAt
func (_m *TokenManager) ValidateAccessToken(createdAt time.Time) bool {
	ret := _m.Called(createdAt)

	var r0 bool
	if rf, ok := ret.Get(0).(func(time.Time) bool); ok {
		r0 = rf(createdAt)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ValidateRefreshToken provides a mock function with given fields: createdAt
func (_m *TokenManager) ValidateRefreshToken(createdAt time.Time) bool {
	ret := _m.Called(createdAt)

	var r0 bool
	if rf, ok := ret.Get(0).(func(time.Time) bool); ok {
		r0 = rf(createdAt)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
