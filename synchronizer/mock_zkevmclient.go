// Code generated by mockery v2.32.0. DO NOT EDIT.

package synchronizer

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	types "github.com/0xPolygonHermez/zkevm-node/jsonrpc/types"
)

// zkEVMClientMock is an autogenerated mock type for the zkEVMClientInterface type
type zkEVMClientMock struct {
	mock.Mock
}

// BlockByNumber provides a mock function with given fields: ctx, number
func (_m *zkEVMClientMock) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	ret := _m.Called(ctx, number)

	var r0 *types.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) (*types.Block, error)); ok {
		return rf(ctx, number)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Block); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockNumber provides a mock function with given fields: ctx
func (_m *zkEVMClientMock) BlockNumber(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExitRootsByGER provides a mock function with given fields: ctx, globalExitRoot
func (_m *zkEVMClientMock) ExitRootsByGER(ctx context.Context, globalExitRoot common.Hash) (*types.ExitRoots, error) {
	ret := _m.Called(ctx, globalExitRoot)

	var r0 *types.ExitRoots
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (*types.ExitRoots, error)); ok {
		return rf(ctx, globalExitRoot)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.ExitRoots); ok {
		r0 = rf(ctx, globalExitRoot)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ExitRoots)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, globalExitRoot)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newZkEVMClientMock creates a new instance of zkEVMClientMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newZkEVMClientMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *zkEVMClientMock {
	mock := &zkEVMClientMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
