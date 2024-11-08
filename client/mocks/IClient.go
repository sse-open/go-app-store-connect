// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"

	request "github.com/sse-open/go-app-store-connect/client/request"
	mock "github.com/stretchr/testify/mock"

	response "github.com/sse-open/go-app-store-connect/client/response"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

type IClient_Expecter struct {
	mock *mock.Mock
}

func (_m *IClient) EXPECT() *IClient_Expecter {
	return &IClient_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: ctx, path
func (_m *IClient) Delete(ctx context.Context, path string) (*response.ClientResponse, error) {
	ret := _m.Called(ctx, path)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 *response.ClientResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*response.ClientResponse, error)); ok {
		return rf(ctx, path)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *response.ClientResponse); ok {
		r0 = rf(ctx, path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*response.ClientResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IClient_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type IClient_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - path string
func (_e *IClient_Expecter) Delete(ctx interface{}, path interface{}) *IClient_Delete_Call {
	return &IClient_Delete_Call{Call: _e.mock.On("Delete", ctx, path)}
}

func (_c *IClient_Delete_Call) Run(run func(ctx context.Context, path string)) *IClient_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *IClient_Delete_Call) Return(_a0 *response.ClientResponse, _a1 error) *IClient_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IClient_Delete_Call) RunAndReturn(run func(context.Context, string) (*response.ClientResponse, error)) *IClient_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, path, query, respPayload
func (_m *IClient) Get(ctx context.Context, path string, query interface{}, respPayload interface{}) (*response.ClientResponse, error) {
	ret := _m.Called(ctx, path, query, respPayload)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *response.ClientResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, interface{}) (*response.ClientResponse, error)); ok {
		return rf(ctx, path, query, respPayload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, interface{}) *response.ClientResponse); ok {
		r0 = rf(ctx, path, query, respPayload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*response.ClientResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}, interface{}) error); ok {
		r1 = rf(ctx, path, query, respPayload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IClient_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type IClient_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - path string
//   - query interface{}
//   - respPayload interface{}
func (_e *IClient_Expecter) Get(ctx interface{}, path interface{}, query interface{}, respPayload interface{}) *IClient_Get_Call {
	return &IClient_Get_Call{Call: _e.mock.On("Get", ctx, path, query, respPayload)}
}

func (_c *IClient_Get_Call) Run(run func(ctx context.Context, path string, query interface{}, respPayload interface{})) *IClient_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}), args[3].(interface{}))
	})
	return _c
}

func (_c *IClient_Get_Call) Return(_a0 *response.ClientResponse, _a1 error) *IClient_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IClient_Get_Call) RunAndReturn(run func(context.Context, string, interface{}, interface{}) (*response.ClientResponse, error)) *IClient_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, path, body, respPayload
func (_m *IClient) Patch(ctx context.Context, path string, body *request.AppStoreConnectRequestPayload, respPayload interface{}) (*response.ClientResponse, error) {
	ret := _m.Called(ctx, path, body, respPayload)

	if len(ret) == 0 {
		panic("no return value specified for Patch")
	}

	var r0 *response.ClientResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *request.AppStoreConnectRequestPayload, interface{}) (*response.ClientResponse, error)); ok {
		return rf(ctx, path, body, respPayload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *request.AppStoreConnectRequestPayload, interface{}) *response.ClientResponse); ok {
		r0 = rf(ctx, path, body, respPayload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*response.ClientResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *request.AppStoreConnectRequestPayload, interface{}) error); ok {
		r1 = rf(ctx, path, body, respPayload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IClient_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type IClient_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - path string
//   - body *request.AppStoreConnectRequestPayload
//   - respPayload interface{}
func (_e *IClient_Expecter) Patch(ctx interface{}, path interface{}, body interface{}, respPayload interface{}) *IClient_Patch_Call {
	return &IClient_Patch_Call{Call: _e.mock.On("Patch", ctx, path, body, respPayload)}
}

func (_c *IClient_Patch_Call) Run(run func(ctx context.Context, path string, body *request.AppStoreConnectRequestPayload, respPayload interface{})) *IClient_Patch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*request.AppStoreConnectRequestPayload), args[3].(interface{}))
	})
	return _c
}

func (_c *IClient_Patch_Call) Return(_a0 *response.ClientResponse, _a1 error) *IClient_Patch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IClient_Patch_Call) RunAndReturn(run func(context.Context, string, *request.AppStoreConnectRequestPayload, interface{}) (*response.ClientResponse, error)) *IClient_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// Post provides a mock function with given fields: ctx, path, body, respPayload
func (_m *IClient) Post(ctx context.Context, path string, body *request.AppStoreConnectRequestPayload, respPayload interface{}) (*response.ClientResponse, error) {
	ret := _m.Called(ctx, path, body, respPayload)

	if len(ret) == 0 {
		panic("no return value specified for Post")
	}

	var r0 *response.ClientResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *request.AppStoreConnectRequestPayload, interface{}) (*response.ClientResponse, error)); ok {
		return rf(ctx, path, body, respPayload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *request.AppStoreConnectRequestPayload, interface{}) *response.ClientResponse); ok {
		r0 = rf(ctx, path, body, respPayload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*response.ClientResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *request.AppStoreConnectRequestPayload, interface{}) error); ok {
		r1 = rf(ctx, path, body, respPayload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IClient_Post_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Post'
type IClient_Post_Call struct {
	*mock.Call
}

// Post is a helper method to define mock.On call
//   - ctx context.Context
//   - path string
//   - body *request.AppStoreConnectRequestPayload
//   - respPayload interface{}
func (_e *IClient_Expecter) Post(ctx interface{}, path interface{}, body interface{}, respPayload interface{}) *IClient_Post_Call {
	return &IClient_Post_Call{Call: _e.mock.On("Post", ctx, path, body, respPayload)}
}

func (_c *IClient_Post_Call) Run(run func(ctx context.Context, path string, body *request.AppStoreConnectRequestPayload, respPayload interface{})) *IClient_Post_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*request.AppStoreConnectRequestPayload), args[3].(interface{}))
	})
	return _c
}

func (_c *IClient_Post_Call) Return(_a0 *response.ClientResponse, _a1 error) *IClient_Post_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IClient_Post_Call) RunAndReturn(run func(context.Context, string, *request.AppStoreConnectRequestPayload, interface{}) (*response.ClientResponse, error)) *IClient_Post_Call {
	_c.Call.Return(run)
	return _c
}

// SetBaseURL provides a mock function with given fields: baseURL
func (_m *IClient) SetBaseURL(baseURL string) {
	_m.Called(baseURL)
}

// IClient_SetBaseURL_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetBaseURL'
type IClient_SetBaseURL_Call struct {
	*mock.Call
}

// SetBaseURL is a helper method to define mock.On call
//   - baseURL string
func (_e *IClient_Expecter) SetBaseURL(baseURL interface{}) *IClient_SetBaseURL_Call {
	return &IClient_SetBaseURL_Call{Call: _e.mock.On("SetBaseURL", baseURL)}
}

func (_c *IClient_SetBaseURL_Call) Run(run func(baseURL string)) *IClient_SetBaseURL_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *IClient_SetBaseURL_Call) Return() *IClient_SetBaseURL_Call {
	_c.Call.Return()
	return _c
}

func (_c *IClient_SetBaseURL_Call) RunAndReturn(run func(string)) *IClient_SetBaseURL_Call {
	_c.Call.Return(run)
	return _c
}

// NewIClient creates a new instance of IClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *IClient {
	mock := &IClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
