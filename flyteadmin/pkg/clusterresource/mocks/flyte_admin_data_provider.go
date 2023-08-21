// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	admin "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/admin"

	mock "github.com/stretchr/testify/mock"
)

// FlyteAdminDataProvider is an autogenerated mock type for the FlyteAdminDataProvider type
type FlyteAdminDataProvider struct {
	mock.Mock
}

type FlyteAdminDataProvider_GetClusterResourceAttributes struct {
	*mock.Call
}

func (_m FlyteAdminDataProvider_GetClusterResourceAttributes) Return(_a0 *admin.ClusterResourceAttributes, _a1 error) *FlyteAdminDataProvider_GetClusterResourceAttributes {
	return &FlyteAdminDataProvider_GetClusterResourceAttributes{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *FlyteAdminDataProvider) OnGetClusterResourceAttributes(ctx context.Context, project string, domain string) *FlyteAdminDataProvider_GetClusterResourceAttributes {
	c_call := _m.On("GetClusterResourceAttributes", ctx, project, domain)
	return &FlyteAdminDataProvider_GetClusterResourceAttributes{Call: c_call}
}

func (_m *FlyteAdminDataProvider) OnGetClusterResourceAttributesMatch(matchers ...interface{}) *FlyteAdminDataProvider_GetClusterResourceAttributes {
	c_call := _m.On("GetClusterResourceAttributes", matchers...)
	return &FlyteAdminDataProvider_GetClusterResourceAttributes{Call: c_call}
}

// GetClusterResourceAttributes provides a mock function with given fields: ctx, project, domain
func (_m *FlyteAdminDataProvider) GetClusterResourceAttributes(ctx context.Context, project string, domain string) (*admin.ClusterResourceAttributes, error) {
	ret := _m.Called(ctx, project, domain)

	var r0 *admin.ClusterResourceAttributes
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *admin.ClusterResourceAttributes); ok {
		r0 = rf(ctx, project, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.ClusterResourceAttributes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, project, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type FlyteAdminDataProvider_GetProjects struct {
	*mock.Call
}

func (_m FlyteAdminDataProvider_GetProjects) Return(_a0 *admin.Projects, _a1 error) *FlyteAdminDataProvider_GetProjects {
	return &FlyteAdminDataProvider_GetProjects{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *FlyteAdminDataProvider) OnGetProjects(ctx context.Context) *FlyteAdminDataProvider_GetProjects {
	c_call := _m.On("GetProjects", ctx)
	return &FlyteAdminDataProvider_GetProjects{Call: c_call}
}

func (_m *FlyteAdminDataProvider) OnGetProjectsMatch(matchers ...interface{}) *FlyteAdminDataProvider_GetProjects {
	c_call := _m.On("GetProjects", matchers...)
	return &FlyteAdminDataProvider_GetProjects{Call: c_call}
}

// GetProjects provides a mock function with given fields: ctx
func (_m *FlyteAdminDataProvider) GetProjects(ctx context.Context) (*admin.Projects, error) {
	ret := _m.Called(ctx)

	var r0 *admin.Projects
	if rf, ok := ret.Get(0).(func(context.Context) *admin.Projects); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Projects)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
