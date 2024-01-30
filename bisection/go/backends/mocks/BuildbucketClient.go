// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	buildbucketpb "go.chromium.org/luci/buildbucket/proto"

	mock "github.com/stretchr/testify/mock"

	swarming "go.chromium.org/luci/common/api/swarming/swarming/v1"
)

// BuildbucketClient is an autogenerated mock type for the BuildbucketClient type
type BuildbucketClient struct {
	mock.Mock
}

// CancelBuild provides a mock function with given fields: ctx, buildID, summary
func (_m *BuildbucketClient) CancelBuild(ctx context.Context, buildID int64, summary string) error {
	ret := _m.Called(ctx, buildID, summary)

	if len(ret) == 0 {
		panic("no return value specified for CancelBuild")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) error); ok {
		r0 = rf(ctx, buildID, summary)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBuildFromWaterfall provides a mock function with given fields: ctx, builderName, commit
func (_m *BuildbucketClient) GetBuildFromWaterfall(ctx context.Context, builderName string, commit string) (*buildbucketpb.Build, error) {
	ret := _m.Called(ctx, builderName, commit)

	if len(ret) == 0 {
		panic("no return value specified for GetBuildFromWaterfall")
	}

	var r0 *buildbucketpb.Build
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*buildbucketpb.Build, error)); ok {
		return rf(ctx, builderName, commit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *buildbucketpb.Build); ok {
		r0 = rf(ctx, builderName, commit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*buildbucketpb.Build)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, builderName, commit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBuildStatus provides a mock function with given fields: ctx, buildID
func (_m *BuildbucketClient) GetBuildStatus(ctx context.Context, buildID int64) (buildbucketpb.Status, error) {
	ret := _m.Called(ctx, buildID)

	if len(ret) == 0 {
		panic("no return value specified for GetBuildStatus")
	}

	var r0 buildbucketpb.Status
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (buildbucketpb.Status, error)); ok {
		return rf(ctx, buildID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) buildbucketpb.Status); ok {
		r0 = rf(ctx, buildID)
	} else {
		r0 = ret.Get(0).(buildbucketpb.Status)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, buildID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBuildWithDeps provides a mock function with given fields: ctx, builderName, bucket, commit, deps
func (_m *BuildbucketClient) GetBuildWithDeps(ctx context.Context, builderName string, bucket string, commit string, deps map[string]interface{}) (*buildbucketpb.Build, error) {
	ret := _m.Called(ctx, builderName, bucket, commit, deps)

	if len(ret) == 0 {
		panic("no return value specified for GetBuildWithDeps")
	}

	var r0 *buildbucketpb.Build
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, map[string]interface{}) (*buildbucketpb.Build, error)); ok {
		return rf(ctx, builderName, bucket, commit, deps)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, map[string]interface{}) *buildbucketpb.Build); ok {
		r0 = rf(ctx, builderName, bucket, commit, deps)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*buildbucketpb.Build)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, map[string]interface{}) error); ok {
		r1 = rf(ctx, builderName, bucket, commit, deps)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBuildWithPatches provides a mock function with given fields: ctx, builderName, bucket, commit, patches
func (_m *BuildbucketClient) GetBuildWithPatches(ctx context.Context, builderName string, bucket string, commit string, patches []*buildbucketpb.GerritChange) (*buildbucketpb.Build, error) {
	ret := _m.Called(ctx, builderName, bucket, commit, patches)

	if len(ret) == 0 {
		panic("no return value specified for GetBuildWithPatches")
	}

	var r0 *buildbucketpb.Build
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, []*buildbucketpb.GerritChange) (*buildbucketpb.Build, error)); ok {
		return rf(ctx, builderName, bucket, commit, patches)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, []*buildbucketpb.GerritChange) *buildbucketpb.Build); ok {
		r0 = rf(ctx, builderName, bucket, commit, patches)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*buildbucketpb.Build)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, []*buildbucketpb.GerritChange) error); ok {
		r1 = rf(ctx, builderName, bucket, commit, patches)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCASReference provides a mock function with given fields: ctx, buildID, target
func (_m *BuildbucketClient) GetCASReference(ctx context.Context, buildID int64, target string) (*swarming.SwarmingRpcsCASReference, error) {
	ret := _m.Called(ctx, buildID, target)

	if len(ret) == 0 {
		panic("no return value specified for GetCASReference")
	}

	var r0 *swarming.SwarmingRpcsCASReference
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) (*swarming.SwarmingRpcsCASReference, error)); ok {
		return rf(ctx, buildID, target)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) *swarming.SwarmingRpcsCASReference); ok {
		r0 = rf(ctx, buildID, target)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swarming.SwarmingRpcsCASReference)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(ctx, buildID, target)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSingleBuild provides a mock function with given fields: ctx, builderName, bucket, commit, deps, patches
func (_m *BuildbucketClient) GetSingleBuild(ctx context.Context, builderName string, bucket string, commit string, deps map[string]interface{}, patches []*buildbucketpb.GerritChange) (*buildbucketpb.Build, error) {
	ret := _m.Called(ctx, builderName, bucket, commit, deps, patches)

	if len(ret) == 0 {
		panic("no return value specified for GetSingleBuild")
	}

	var r0 *buildbucketpb.Build
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, map[string]interface{}, []*buildbucketpb.GerritChange) (*buildbucketpb.Build, error)); ok {
		return rf(ctx, builderName, bucket, commit, deps, patches)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, map[string]interface{}, []*buildbucketpb.GerritChange) *buildbucketpb.Build); ok {
		r0 = rf(ctx, builderName, bucket, commit, deps, patches)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*buildbucketpb.Build)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, map[string]interface{}, []*buildbucketpb.GerritChange) error); ok {
		r1 = rf(ctx, builderName, bucket, commit, deps, patches)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartChromeBuild provides a mock function with given fields: ctx, pinpointJobID, requestID, builderName, commitHash, deps, patches
func (_m *BuildbucketClient) StartChromeBuild(ctx context.Context, pinpointJobID string, requestID string, builderName string, commitHash string, deps map[string]interface{}, patches []*buildbucketpb.GerritChange) (*buildbucketpb.Build, error) {
	ret := _m.Called(ctx, pinpointJobID, requestID, builderName, commitHash, deps, patches)

	if len(ret) == 0 {
		panic("no return value specified for StartChromeBuild")
	}

	var r0 *buildbucketpb.Build
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, map[string]interface{}, []*buildbucketpb.GerritChange) (*buildbucketpb.Build, error)); ok {
		return rf(ctx, pinpointJobID, requestID, builderName, commitHash, deps, patches)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, map[string]interface{}, []*buildbucketpb.GerritChange) *buildbucketpb.Build); ok {
		r0 = rf(ctx, pinpointJobID, requestID, builderName, commitHash, deps, patches)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*buildbucketpb.Build)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, map[string]interface{}, []*buildbucketpb.GerritChange) error); ok {
		r1 = rf(ctx, pinpointJobID, requestID, builderName, commitHash, deps, patches)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBuildbucketClient creates a new instance of BuildbucketClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBuildbucketClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *BuildbucketClient {
	mock := &BuildbucketClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
