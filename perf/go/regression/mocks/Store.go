// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	clustering2 "go.skia.org/infra/perf/go/clustering2"

	frame "go.skia.org/infra/perf/go/ui/frame"

	mock "github.com/stretchr/testify/mock"

	regression "go.skia.org/infra/perf/go/regression"

	types "go.skia.org/infra/perf/go/types"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// GetByIDs provides a mock function with given fields: ctx, ids
func (_m *Store) GetByIDs(ctx context.Context, ids []string) ([]*regression.Regression, error) {
	ret := _m.Called(ctx, ids)

	if len(ret) == 0 {
		panic("no return value specified for GetByIDs")
	}

	var r0 []*regression.Regression
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) ([]*regression.Regression, error)); ok {
		return rf(ctx, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) []*regression.Regression); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*regression.Regression)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRegressionsBySubName provides a mock function with given fields: ctx, sub_name, limit, offset
func (_m *Store) GetRegressionsBySubName(ctx context.Context, sub_name string, limit int, offset int) ([]*regression.Regression, error) {
	ret := _m.Called(ctx, sub_name, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetRegressionsBySubName")
	}

	var r0 []*regression.Regression
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) ([]*regression.Regression, error)); ok {
		return rf(ctx, sub_name, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) []*regression.Regression); ok {
		r0 = rf(ctx, sub_name, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*regression.Regression)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) error); ok {
		r1 = rf(ctx, sub_name, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Range provides a mock function with given fields: ctx, begin, end
func (_m *Store) Range(ctx context.Context, begin types.CommitNumber, end types.CommitNumber) (map[types.CommitNumber]*regression.AllRegressionsForCommit, error) {
	ret := _m.Called(ctx, begin, end)

	if len(ret) == 0 {
		panic("no return value specified for Range")
	}

	var r0 map[types.CommitNumber]*regression.AllRegressionsForCommit
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.CommitNumber, types.CommitNumber) (map[types.CommitNumber]*regression.AllRegressionsForCommit, error)); ok {
		return rf(ctx, begin, end)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.CommitNumber, types.CommitNumber) map[types.CommitNumber]*regression.AllRegressionsForCommit); ok {
		r0 = rf(ctx, begin, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[types.CommitNumber]*regression.AllRegressionsForCommit)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.CommitNumber, types.CommitNumber) error); ok {
		r1 = rf(ctx, begin, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetHigh provides a mock function with given fields: ctx, commitNumber, alertID, df, high
func (_m *Store) SetHigh(ctx context.Context, commitNumber types.CommitNumber, alertID string, df *frame.FrameResponse, high *clustering2.ClusterSummary) (bool, string, error) {
	ret := _m.Called(ctx, commitNumber, alertID, df, high)

	if len(ret) == 0 {
		panic("no return value specified for SetHigh")
	}

	var r0 bool
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, types.CommitNumber, string, *frame.FrameResponse, *clustering2.ClusterSummary) (bool, string, error)); ok {
		return rf(ctx, commitNumber, alertID, df, high)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.CommitNumber, string, *frame.FrameResponse, *clustering2.ClusterSummary) bool); ok {
		r0 = rf(ctx, commitNumber, alertID, df, high)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.CommitNumber, string, *frame.FrameResponse, *clustering2.ClusterSummary) string); ok {
		r1 = rf(ctx, commitNumber, alertID, df, high)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, types.CommitNumber, string, *frame.FrameResponse, *clustering2.ClusterSummary) error); ok {
		r2 = rf(ctx, commitNumber, alertID, df, high)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SetLow provides a mock function with given fields: ctx, commitNumber, alertID, df, low
func (_m *Store) SetLow(ctx context.Context, commitNumber types.CommitNumber, alertID string, df *frame.FrameResponse, low *clustering2.ClusterSummary) (bool, string, error) {
	ret := _m.Called(ctx, commitNumber, alertID, df, low)

	if len(ret) == 0 {
		panic("no return value specified for SetLow")
	}

	var r0 bool
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, types.CommitNumber, string, *frame.FrameResponse, *clustering2.ClusterSummary) (bool, string, error)); ok {
		return rf(ctx, commitNumber, alertID, df, low)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.CommitNumber, string, *frame.FrameResponse, *clustering2.ClusterSummary) bool); ok {
		r0 = rf(ctx, commitNumber, alertID, df, low)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.CommitNumber, string, *frame.FrameResponse, *clustering2.ClusterSummary) string); ok {
		r1 = rf(ctx, commitNumber, alertID, df, low)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, types.CommitNumber, string, *frame.FrameResponse, *clustering2.ClusterSummary) error); ok {
		r2 = rf(ctx, commitNumber, alertID, df, low)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TriageHigh provides a mock function with given fields: ctx, commitNumber, alertID, tr
func (_m *Store) TriageHigh(ctx context.Context, commitNumber types.CommitNumber, alertID string, tr regression.TriageStatus) error {
	ret := _m.Called(ctx, commitNumber, alertID, tr)

	if len(ret) == 0 {
		panic("no return value specified for TriageHigh")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.CommitNumber, string, regression.TriageStatus) error); ok {
		r0 = rf(ctx, commitNumber, alertID, tr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TriageLow provides a mock function with given fields: ctx, commitNumber, alertID, tr
func (_m *Store) TriageLow(ctx context.Context, commitNumber types.CommitNumber, alertID string, tr regression.TriageStatus) error {
	ret := _m.Called(ctx, commitNumber, alertID, tr)

	if len(ret) == 0 {
		panic("no return value specified for TriageLow")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.CommitNumber, string, regression.TriageStatus) error); ok {
		r0 = rf(ctx, commitNumber, alertID, tr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Write provides a mock function with given fields: ctx, regressions
func (_m *Store) Write(ctx context.Context, regressions map[types.CommitNumber]*regression.AllRegressionsForCommit) error {
	ret := _m.Called(ctx, regressions)

	if len(ret) == 0 {
		panic("no return value specified for Write")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, map[types.CommitNumber]*regression.AllRegressionsForCommit) error); ok {
		r0 = rf(ctx, regressions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *Store {
	mock := &Store{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
