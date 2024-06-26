// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	backends "go.skia.org/infra/pinpoint/go/backends"

	mock "github.com/stretchr/testify/mock"
)

// CrrevClient is an autogenerated mock type for the CrrevClient type
type CrrevClient struct {
	mock.Mock
}

// GetCommitInfo provides a mock function with given fields: ctx, commit
func (_m *CrrevClient) GetCommitInfo(ctx context.Context, commit string) (*backends.CrrevResponse, error) {
	ret := _m.Called(ctx, commit)

	if len(ret) == 0 {
		panic("no return value specified for GetCommitInfo")
	}

	var r0 *backends.CrrevResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*backends.CrrevResponse, error)); ok {
		return rf(ctx, commit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *backends.CrrevResponse); ok {
		r0 = rf(ctx, commit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backends.CrrevResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, commit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCrrevClient creates a new instance of CrrevClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCrrevClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *CrrevClient {
	mock := &CrrevClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
