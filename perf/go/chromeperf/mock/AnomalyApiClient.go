// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	chromeperf "go.skia.org/infra/perf/go/chromeperf"

	mock "github.com/stretchr/testify/mock"
)

// AnomalyApiClient is an autogenerated mock type for the AnomalyApiClient type
type AnomalyApiClient struct {
	mock.Mock
}

// GetAnomalies provides a mock function with given fields: ctx, traceNames, startCommitPosition, endCommitPosition
func (_m *AnomalyApiClient) GetAnomalies(ctx context.Context, traceNames []string, startCommitPosition int, endCommitPosition int) (chromeperf.AnomalyMap, error) {
	ret := _m.Called(ctx, traceNames, startCommitPosition, endCommitPosition)

	if len(ret) == 0 {
		panic("no return value specified for GetAnomalies")
	}

	var r0 chromeperf.AnomalyMap
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, int, int) (chromeperf.AnomalyMap, error)); ok {
		return rf(ctx, traceNames, startCommitPosition, endCommitPosition)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string, int, int) chromeperf.AnomalyMap); ok {
		r0 = rf(ctx, traceNames, startCommitPosition, endCommitPosition)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chromeperf.AnomalyMap)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string, int, int) error); ok {
		r1 = rf(ctx, traceNames, startCommitPosition, endCommitPosition)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAnomaliesAroundRevision provides a mock function with given fields: ctx, revision
func (_m *AnomalyApiClient) GetAnomaliesAroundRevision(ctx context.Context, revision int) ([]chromeperf.AnomalyForRevision, error) {
	ret := _m.Called(ctx, revision)

	if len(ret) == 0 {
		panic("no return value specified for GetAnomaliesAroundRevision")
	}

	var r0 []chromeperf.AnomalyForRevision
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]chromeperf.AnomalyForRevision, error)); ok {
		return rf(ctx, revision)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []chromeperf.AnomalyForRevision); ok {
		r0 = rf(ctx, revision)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chromeperf.AnomalyForRevision)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, revision)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReportRegression provides a mock function with given fields: ctx, testPath, startCommitPosition, endCommitPosition, projectId, isImprovement, botName, internal, medianBefore, medianAfter
func (_m *AnomalyApiClient) ReportRegression(ctx context.Context, testPath string, startCommitPosition int32, endCommitPosition int32, projectId string, isImprovement bool, botName string, internal bool, medianBefore float32, medianAfter float32) (*chromeperf.ReportRegressionResponse, error) {
	ret := _m.Called(ctx, testPath, startCommitPosition, endCommitPosition, projectId, isImprovement, botName, internal, medianBefore, medianAfter)

	if len(ret) == 0 {
		panic("no return value specified for ReportRegression")
	}

	var r0 *chromeperf.ReportRegressionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int32, int32, string, bool, string, bool, float32, float32) (*chromeperf.ReportRegressionResponse, error)); ok {
		return rf(ctx, testPath, startCommitPosition, endCommitPosition, projectId, isImprovement, botName, internal, medianBefore, medianAfter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int32, int32, string, bool, string, bool, float32, float32) *chromeperf.ReportRegressionResponse); ok {
		r0 = rf(ctx, testPath, startCommitPosition, endCommitPosition, projectId, isImprovement, botName, internal, medianBefore, medianAfter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chromeperf.ReportRegressionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int32, int32, string, bool, string, bool, float32, float32) error); ok {
		r1 = rf(ctx, testPath, startCommitPosition, endCommitPosition, projectId, isImprovement, botName, internal, medianBefore, medianAfter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAnomalyApiClient creates a new instance of AnomalyApiClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAnomalyApiClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *AnomalyApiClient {
	mock := &AnomalyApiClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}