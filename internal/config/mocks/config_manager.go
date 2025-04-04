// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ConfigManager is an autogenerated mock type for the ConfigManager type
type ConfigManager struct {
	mock.Mock
}

// AddFile provides a mock function with given fields: filename
func (_m *ConfigManager) AddFile(filename string) error {
	ret := _m.Called(filename)

	if len(ret) == 0 {
		panic("no return value specified for AddFile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(filename)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with no fields
func (_m *ConfigManager) Save() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewConfigManager creates a new instance of ConfigManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConfigManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConfigManager {
	mock := &ConfigManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
