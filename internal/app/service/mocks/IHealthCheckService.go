// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import (
	context "context"

	response "github.com/vinsensiussatya/bego-training/internal/pkg/response"
	mock "github.com/stretchr/testify/mock"
)

// IHealthCheckService is an autogenerated mock type for the IHealthCheckService type
type IHealthCheckService struct {
	mock.Mock
}

// Ping provides a mock function with given fields: ctx
func (_m *IHealthCheckService) Ping(ctx context.Context) response.PingResponse {
	ret := _m.Called(ctx)

	var r0 response.PingResponse
	if rf, ok := ret.Get(0).(func(context.Context) response.PingResponse); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(response.PingResponse)
	}

	return r0
}

// NewIHealthCheckService creates a new instance of IHealthCheckService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIHealthCheckService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IHealthCheckService {
	mock := &IHealthCheckService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}