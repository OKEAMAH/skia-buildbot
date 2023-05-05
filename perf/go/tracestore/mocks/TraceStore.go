// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	paramtools "go.skia.org/infra/go/paramtools"

	query "go.skia.org/infra/go/query"

	testing "testing"

	time "time"

	tracestore "go.skia.org/infra/perf/go/tracestore"

	types "go.skia.org/infra/perf/go/types"
)

// TraceStore is an autogenerated mock type for the TraceStore type
type TraceStore struct {
	mock.Mock
}

// CommitNumberOfTileStart provides a mock function with given fields: ctx, commitNumber
func (_m *TraceStore) CommitNumberOfTileStart(ctx context.Context, commitNumber types.CommitNumber) (types.CommitNumber, error) {
	ret := _m.Called(ctx, commitNumber)

	var r0 types.CommitNumber
	if rf, ok := ret.Get(0).(func(context.Context, types.CommitNumber) types.CommitNumber); ok {
		r0 = rf(ctx, commitNumber)
	} else {
		r0 = ret.Get(0).(types.CommitNumber)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.CommitNumber) error); ok {
		r1 = rf(ctx, commitNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastNSources provides a mock function with given fields: ctx, traceID, n
func (_m *TraceStore) GetLastNSources(ctx context.Context, traceID string, n int) ([]tracestore.Source, error) {
	ret := _m.Called(ctx, traceID, n)

	var r0 []tracestore.Source
	if rf, ok := ret.Get(0).(func(context.Context, string, int) []tracestore.Source); ok {
		r0 = rf(ctx, traceID, n)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]tracestore.Source)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int) error); ok {
		r1 = rf(ctx, traceID, n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestTile provides a mock function with given fields: _a0
func (_m *TraceStore) GetLatestTile(_a0 context.Context) (types.TileNumber, error) {
	ret := _m.Called(_a0)

	var r0 types.TileNumber
	if rf, ok := ret.Get(0).(func(context.Context) types.TileNumber); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(types.TileNumber)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetParamSet provides a mock function with given fields: ctx, tileNumber
func (_m *TraceStore) GetParamSet(ctx context.Context, tileNumber types.TileNumber) (paramtools.ReadOnlyParamSet, error) {
	ret := _m.Called(ctx, tileNumber)

	var r0 paramtools.ReadOnlyParamSet
	if rf, ok := ret.Get(0).(func(context.Context, types.TileNumber) paramtools.ReadOnlyParamSet); ok {
		r0 = rf(ctx, tileNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(paramtools.ReadOnlyParamSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.TileNumber) error); ok {
		r1 = rf(ctx, tileNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSource provides a mock function with given fields: ctx, commitNumber, traceId
func (_m *TraceStore) GetSource(ctx context.Context, commitNumber types.CommitNumber, traceId string) (string, error) {
	ret := _m.Called(ctx, commitNumber, traceId)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, types.CommitNumber, string) string); ok {
		r0 = rf(ctx, commitNumber, traceId)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.CommitNumber, string) error); ok {
		r1 = rf(ctx, commitNumber, traceId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTraceIDsBySource provides a mock function with given fields: ctx, sourceFilename, tileNumber
func (_m *TraceStore) GetTraceIDsBySource(ctx context.Context, sourceFilename string, tileNumber types.TileNumber) ([]string, error) {
	ret := _m.Called(ctx, sourceFilename, tileNumber)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, string, types.TileNumber) []string); ok {
		r0 = rf(ctx, sourceFilename, tileNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.TileNumber) error); ok {
		r1 = rf(ctx, sourceFilename, tileNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OffsetFromCommitNumber provides a mock function with given fields: commitNumber
func (_m *TraceStore) OffsetFromCommitNumber(commitNumber types.CommitNumber) int32 {
	ret := _m.Called(commitNumber)

	var r0 int32
	if rf, ok := ret.Get(0).(func(types.CommitNumber) int32); ok {
		r0 = rf(commitNumber)
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// QueryTraces provides a mock function with given fields: ctx, tileNumber, q
func (_m *TraceStore) QueryTraces(ctx context.Context, tileNumber types.TileNumber, q *query.Query) (types.TraceSet, error) {
	ret := _m.Called(ctx, tileNumber, q)

	var r0 types.TraceSet
	if rf, ok := ret.Get(0).(func(context.Context, types.TileNumber, *query.Query) types.TraceSet); ok {
		r0 = rf(ctx, tileNumber, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.TraceSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.TileNumber, *query.Query) error); ok {
		r1 = rf(ctx, tileNumber, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryTracesIDOnly provides a mock function with given fields: ctx, tileNumber, q
func (_m *TraceStore) QueryTracesIDOnly(ctx context.Context, tileNumber types.TileNumber, q *query.Query) (<-chan paramtools.Params, error) {
	ret := _m.Called(ctx, tileNumber, q)

	var r0 <-chan paramtools.Params
	if rf, ok := ret.Get(0).(func(context.Context, types.TileNumber, *query.Query) <-chan paramtools.Params); ok {
		r0 = rf(ctx, tileNumber, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan paramtools.Params)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.TileNumber, *query.Query) error); ok {
		r1 = rf(ctx, tileNumber, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadTraces provides a mock function with given fields: ctx, tileNumber, keys
func (_m *TraceStore) ReadTraces(ctx context.Context, tileNumber types.TileNumber, keys []string) (types.TraceSet, error) {
	ret := _m.Called(ctx, tileNumber, keys)

	var r0 types.TraceSet
	if rf, ok := ret.Get(0).(func(context.Context, types.TileNumber, []string) types.TraceSet); ok {
		r0 = rf(ctx, tileNumber, keys)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.TraceSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.TileNumber, []string) error); ok {
		r1 = rf(ctx, tileNumber, keys)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadTracesForCommitRange provides a mock function with given fields: ctx, keys, begin, end
func (_m *TraceStore) ReadTracesForCommitRange(ctx context.Context, keys []string, begin types.CommitNumber, end types.CommitNumber) (types.TraceSet, error) {
	ret := _m.Called(ctx, keys, begin, end)

	var r0 types.TraceSet
	if rf, ok := ret.Get(0).(func(context.Context, []string, types.CommitNumber, types.CommitNumber) types.TraceSet); ok {
		r0 = rf(ctx, keys, begin, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.TraceSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string, types.CommitNumber, types.CommitNumber) error); ok {
		r1 = rf(ctx, keys, begin, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TileNumber provides a mock function with given fields: commitNumber
func (_m *TraceStore) TileNumber(commitNumber types.CommitNumber) types.TileNumber {
	ret := _m.Called(commitNumber)

	var r0 types.TileNumber
	if rf, ok := ret.Get(0).(func(types.CommitNumber) types.TileNumber); ok {
		r0 = rf(commitNumber)
	} else {
		r0 = ret.Get(0).(types.TileNumber)
	}

	return r0
}

// TileSize provides a mock function with given fields:
func (_m *TraceStore) TileSize() int32 {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// TraceCount provides a mock function with given fields: ctx, tileNumber
func (_m *TraceStore) TraceCount(ctx context.Context, tileNumber types.TileNumber) (int64, error) {
	ret := _m.Called(ctx, tileNumber)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, types.TileNumber) int64); ok {
		r0 = rf(ctx, tileNumber)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.TileNumber) error); ok {
		r1 = rf(ctx, tileNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WriteTraces provides a mock function with given fields: ctx, commitNumber, params, values, paramset, source, timestamp
func (_m *TraceStore) WriteTraces(ctx context.Context, commitNumber types.CommitNumber, params []paramtools.Params, values []float32, paramset paramtools.ParamSet, source string, timestamp time.Time) error {
	ret := _m.Called(ctx, commitNumber, params, values, paramset, source, timestamp)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.CommitNumber, []paramtools.Params, []float32, paramtools.ParamSet, string, time.Time) error); ok {
		r0 = rf(ctx, commitNumber, params, values, paramset, source, timestamp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTraceStore creates a new instance of TraceStore. It also registers a cleanup function to assert the mocks expectations.
func NewTraceStore(t testing.TB) *TraceStore {
	mock := &TraceStore{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
