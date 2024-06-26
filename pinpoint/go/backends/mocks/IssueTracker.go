// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	pinpointpb "go.skia.org/infra/pinpoint/proto/v1"
)

// IssueTracker is an autogenerated mock type for the IssueTracker type
type IssueTracker struct {
	mock.Mock
}

// ReportCulprit provides a mock function with given fields: issueID, culprits
func (_m *IssueTracker) ReportCulprit(issueID int64, culprits []*pinpointpb.CombinedCommit) error {
	ret := _m.Called(issueID, culprits)

	if len(ret) == 0 {
		panic("no return value specified for ReportCulprit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, []*pinpointpb.CombinedCommit) error); ok {
		r0 = rf(issueID, culprits)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIssueTracker creates a new instance of IssueTracker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIssueTracker(t interface {
	mock.TestingT
	Cleanup(func())
}) *IssueTracker {
	mock := &IssueTracker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
