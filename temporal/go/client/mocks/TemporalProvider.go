// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	client "go.temporal.io/sdk/client"

	mock "github.com/stretchr/testify/mock"
)

// TemporalProvider is an autogenerated mock type for the TemporalProvider type
type TemporalProvider struct {
	mock.Mock
}

// NewClient provides a mock function with given fields: _a0, _a1
func (_m *TemporalProvider) NewClient(_a0 string, _a1 string) (client.Client, func(), error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for NewClient")
	}

	var r0 client.Client
	var r1 func()
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (client.Client, func(), error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(string, string) client.Client); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.Client)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) func()); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(func())
		}
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewTemporalProvider creates a new instance of TemporalProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTemporalProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *TemporalProvider {
	mock := &TemporalProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}