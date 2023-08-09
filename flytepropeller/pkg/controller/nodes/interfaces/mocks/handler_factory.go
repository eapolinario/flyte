// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	interfaces "github.com/flyteorg/flytepropeller/pkg/controller/nodes/interfaces"
	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/flyteorg/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

// HandlerFactory is an autogenerated mock type for the HandlerFactory type
type HandlerFactory struct {
	mock.Mock
}

type HandlerFactory_GetHandler struct {
	*mock.Call
}

func (_m HandlerFactory_GetHandler) Return(_a0 interfaces.NodeHandler, _a1 error) *HandlerFactory_GetHandler {
	return &HandlerFactory_GetHandler{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *HandlerFactory) OnGetHandler(kind v1alpha1.NodeKind) *HandlerFactory_GetHandler {
	c_call := _m.On("GetHandler", kind)
	return &HandlerFactory_GetHandler{Call: c_call}
}

func (_m *HandlerFactory) OnGetHandlerMatch(matchers ...interface{}) *HandlerFactory_GetHandler {
	c_call := _m.On("GetHandler", matchers...)
	return &HandlerFactory_GetHandler{Call: c_call}
}

// GetHandler provides a mock function with given fields: kind
func (_m *HandlerFactory) GetHandler(kind v1alpha1.NodeKind) (interfaces.NodeHandler, error) {
	ret := _m.Called(kind)

	var r0 interfaces.NodeHandler
	if rf, ok := ret.Get(0).(func(v1alpha1.NodeKind) interfaces.NodeHandler); ok {
		r0 = rf(kind)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.NodeHandler)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(v1alpha1.NodeKind) error); ok {
		r1 = rf(kind)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type HandlerFactory_Setup struct {
	*mock.Call
}

func (_m HandlerFactory_Setup) Return(_a0 error) *HandlerFactory_Setup {
	return &HandlerFactory_Setup{Call: _m.Call.Return(_a0)}
}

func (_m *HandlerFactory) OnSetup(ctx context.Context, executor interfaces.Node, setup interfaces.SetupContext) *HandlerFactory_Setup {
	c_call := _m.On("Setup", ctx, executor, setup)
	return &HandlerFactory_Setup{Call: c_call}
}

func (_m *HandlerFactory) OnSetupMatch(matchers ...interface{}) *HandlerFactory_Setup {
	c_call := _m.On("Setup", matchers...)
	return &HandlerFactory_Setup{Call: c_call}
}

// Setup provides a mock function with given fields: ctx, executor, setup
func (_m *HandlerFactory) Setup(ctx context.Context, executor interfaces.Node, setup interfaces.SetupContext) error {
	ret := _m.Called(ctx, executor, setup)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interfaces.Node, interfaces.SetupContext) error); ok {
		r0 = rf(ctx, executor, setup)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
