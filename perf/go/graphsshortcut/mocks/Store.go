// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	graphsshortcut "go.skia.org/infra/perf/go/graphsshortcut"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// GetShortcut provides a mock function with given fields: ctx, id
func (_m *Store) GetShortcut(ctx context.Context, id string) (*graphsshortcut.GraphsShortcut, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetShortcut")
	}

	var r0 *graphsshortcut.GraphsShortcut
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*graphsshortcut.GraphsShortcut, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *graphsshortcut.GraphsShortcut); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*graphsshortcut.GraphsShortcut)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertShortcut provides a mock function with given fields: ctx, shortcut
func (_m *Store) InsertShortcut(ctx context.Context, shortcut *graphsshortcut.GraphsShortcut) (string, error) {
	ret := _m.Called(ctx, shortcut)

	if len(ret) == 0 {
		panic("no return value specified for InsertShortcut")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *graphsshortcut.GraphsShortcut) (string, error)); ok {
		return rf(ctx, shortcut)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *graphsshortcut.GraphsShortcut) string); ok {
		r0 = rf(ctx, shortcut)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *graphsshortcut.GraphsShortcut) error); ok {
		r1 = rf(ctx, shortcut)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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