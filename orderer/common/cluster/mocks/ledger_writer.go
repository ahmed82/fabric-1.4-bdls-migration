// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import common "github.com/hyperledger/fabric-protos-go/common"
import mock "github.com/stretchr/testify/mock"

// LedgerWriter is an autogenerated mock type for the LedgerWriter type
type LedgerWriter struct {
	mock.Mock
}

// Append provides a mock function with given fields: block
func (_m *LedgerWriter) Append(block *common.Block) error {
	ret := _m.Called(block)

	var r0 error
	if rf, ok := ret.Get(0).(func(*common.Block) error); ok {
		r0 = rf(block)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Height provides a mock function with given fields:
func (_m *LedgerWriter) Height() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}
