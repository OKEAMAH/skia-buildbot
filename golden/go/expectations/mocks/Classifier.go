// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	expectations "go.skia.org/infra/golden/go/expectations"

	types "go.skia.org/infra/golden/go/types"
)

// Classifier is an autogenerated mock type for the Classifier type
type Classifier struct {
	mock.Mock
}

// Classification provides a mock function with given fields: test, digest
func (_m *Classifier) Classification(test types.TestName, digest types.Digest) expectations.LabelStr {
	ret := _m.Called(test, digest)

	var r0 expectations.LabelStr
	if rf, ok := ret.Get(0).(func(types.TestName, types.Digest) expectations.LabelStr); ok {
		r0 = rf(test, digest)
	} else {
		r0 = ret.Get(0).(expectations.LabelStr)
	}

	return r0
}
