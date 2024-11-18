// Code generated by mockery. DO NOT EDIT.

package mockmetastorev1

import (
	context "context"

	grpc "google.golang.org/grpc"

	metastorev1 "github.com/grafana/pyroscope/api/gen/proto/go/metastore/v1"

	mock "github.com/stretchr/testify/mock"
)

// MockMetadataQueryServiceClient is an autogenerated mock type for the MetadataQueryServiceClient type
type MockMetadataQueryServiceClient struct {
	mock.Mock
}

type MockMetadataQueryServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMetadataQueryServiceClient) EXPECT() *MockMetadataQueryServiceClient_Expecter {
	return &MockMetadataQueryServiceClient_Expecter{mock: &_m.Mock}
}

// QueryMetadata provides a mock function with given fields: ctx, in, opts
func (_m *MockMetadataQueryServiceClient) QueryMetadata(ctx context.Context, in *metastorev1.QueryMetadataRequest, opts ...grpc.CallOption) (*metastorev1.QueryMetadataResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryMetadata")
	}

	var r0 *metastorev1.QueryMetadataResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *metastorev1.QueryMetadataRequest, ...grpc.CallOption) (*metastorev1.QueryMetadataResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *metastorev1.QueryMetadataRequest, ...grpc.CallOption) *metastorev1.QueryMetadataResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metastorev1.QueryMetadataResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *metastorev1.QueryMetadataRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMetadataQueryServiceClient_QueryMetadata_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryMetadata'
type MockMetadataQueryServiceClient_QueryMetadata_Call struct {
	*mock.Call
}

// QueryMetadata is a helper method to define mock.On call
//   - ctx context.Context
//   - in *metastorev1.QueryMetadataRequest
//   - opts ...grpc.CallOption
func (_e *MockMetadataQueryServiceClient_Expecter) QueryMetadata(ctx interface{}, in interface{}, opts ...interface{}) *MockMetadataQueryServiceClient_QueryMetadata_Call {
	return &MockMetadataQueryServiceClient_QueryMetadata_Call{Call: _e.mock.On("QueryMetadata",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockMetadataQueryServiceClient_QueryMetadata_Call) Run(run func(ctx context.Context, in *metastorev1.QueryMetadataRequest, opts ...grpc.CallOption)) *MockMetadataQueryServiceClient_QueryMetadata_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*metastorev1.QueryMetadataRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockMetadataQueryServiceClient_QueryMetadata_Call) Return(_a0 *metastorev1.QueryMetadataResponse, _a1 error) *MockMetadataQueryServiceClient_QueryMetadata_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMetadataQueryServiceClient_QueryMetadata_Call) RunAndReturn(run func(context.Context, *metastorev1.QueryMetadataRequest, ...grpc.CallOption) (*metastorev1.QueryMetadataResponse, error)) *MockMetadataQueryServiceClient_QueryMetadata_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMetadataQueryServiceClient creates a new instance of MockMetadataQueryServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMetadataQueryServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMetadataQueryServiceClient {
	mock := &MockMetadataQueryServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}