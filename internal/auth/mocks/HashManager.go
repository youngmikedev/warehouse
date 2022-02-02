// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// HashManager is an autogenerated mock type for the HashManager type
type HashManager struct {
	mock.Mock
}

// Hash provides a mock function with given fields: str
func (_m *HashManager) Hash(str string) (string, error) {
	ret := _m.Called(str)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(str)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(str)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Validate provides a mock function with given fields: hashedPassword, password
func (_m *HashManager) Validate(hashedPassword string, password string) bool {
	ret := _m.Called(hashedPassword, password)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(hashedPassword, password)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
