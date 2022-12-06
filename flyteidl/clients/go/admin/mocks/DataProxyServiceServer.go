// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	service "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/service"
	mock "github.com/stretchr/testify/mock"
)

// DataProxyServiceServer is an autogenerated mock type for the DataProxyServiceServer type
type DataProxyServiceServer struct {
	mock.Mock
}

type DataProxyServiceServer_CreateDownloadLink struct {
	*mock.Call
}

func (_m DataProxyServiceServer_CreateDownloadLink) Return(_a0 *service.CreateDownloadLinkResponse, _a1 error) *DataProxyServiceServer_CreateDownloadLink {
	return &DataProxyServiceServer_CreateDownloadLink{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *DataProxyServiceServer) OnCreateDownloadLink(_a0 context.Context, _a1 *service.CreateDownloadLinkRequest) *DataProxyServiceServer_CreateDownloadLink {
	c_call := _m.On("CreateDownloadLink", _a0, _a1)
	return &DataProxyServiceServer_CreateDownloadLink{Call: c_call}
}

func (_m *DataProxyServiceServer) OnCreateDownloadLinkMatch(matchers ...interface{}) *DataProxyServiceServer_CreateDownloadLink {
	c_call := _m.On("CreateDownloadLink", matchers...)
	return &DataProxyServiceServer_CreateDownloadLink{Call: c_call}
}

// CreateDownloadLink provides a mock function with given fields: _a0, _a1
func (_m *DataProxyServiceServer) CreateDownloadLink(_a0 context.Context, _a1 *service.CreateDownloadLinkRequest) (*service.CreateDownloadLinkResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *service.CreateDownloadLinkResponse
	if rf, ok := ret.Get(0).(func(context.Context, *service.CreateDownloadLinkRequest) *service.CreateDownloadLinkResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.CreateDownloadLinkResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.CreateDownloadLinkRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DataProxyServiceServer_CreateDownloadLocation struct {
	*mock.Call
}

func (_m DataProxyServiceServer_CreateDownloadLocation) Return(_a0 *service.CreateDownloadLocationResponse, _a1 error) *DataProxyServiceServer_CreateDownloadLocation {
	return &DataProxyServiceServer_CreateDownloadLocation{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *DataProxyServiceServer) OnCreateDownloadLocation(_a0 context.Context, _a1 *service.CreateDownloadLocationRequest) *DataProxyServiceServer_CreateDownloadLocation {
	c_call := _m.On("CreateDownloadLocation", _a0, _a1)
	return &DataProxyServiceServer_CreateDownloadLocation{Call: c_call}
}

func (_m *DataProxyServiceServer) OnCreateDownloadLocationMatch(matchers ...interface{}) *DataProxyServiceServer_CreateDownloadLocation {
	c_call := _m.On("CreateDownloadLocation", matchers...)
	return &DataProxyServiceServer_CreateDownloadLocation{Call: c_call}
}

// CreateDownloadLocation provides a mock function with given fields: _a0, _a1
func (_m *DataProxyServiceServer) CreateDownloadLocation(_a0 context.Context, _a1 *service.CreateDownloadLocationRequest) (*service.CreateDownloadLocationResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *service.CreateDownloadLocationResponse
	if rf, ok := ret.Get(0).(func(context.Context, *service.CreateDownloadLocationRequest) *service.CreateDownloadLocationResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.CreateDownloadLocationResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.CreateDownloadLocationRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DataProxyServiceServer_CreateUploadLocation struct {
	*mock.Call
}

func (_m DataProxyServiceServer_CreateUploadLocation) Return(_a0 *service.CreateUploadLocationResponse, _a1 error) *DataProxyServiceServer_CreateUploadLocation {
	return &DataProxyServiceServer_CreateUploadLocation{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *DataProxyServiceServer) OnCreateUploadLocation(_a0 context.Context, _a1 *service.CreateUploadLocationRequest) *DataProxyServiceServer_CreateUploadLocation {
	c_call := _m.On("CreateUploadLocation", _a0, _a1)
	return &DataProxyServiceServer_CreateUploadLocation{Call: c_call}
}

func (_m *DataProxyServiceServer) OnCreateUploadLocationMatch(matchers ...interface{}) *DataProxyServiceServer_CreateUploadLocation {
	c_call := _m.On("CreateUploadLocation", matchers...)
	return &DataProxyServiceServer_CreateUploadLocation{Call: c_call}
}

// CreateUploadLocation provides a mock function with given fields: _a0, _a1
func (_m *DataProxyServiceServer) CreateUploadLocation(_a0 context.Context, _a1 *service.CreateUploadLocationRequest) (*service.CreateUploadLocationResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *service.CreateUploadLocationResponse
	if rf, ok := ret.Get(0).(func(context.Context, *service.CreateUploadLocationRequest) *service.CreateUploadLocationResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.CreateUploadLocationResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.CreateUploadLocationRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
