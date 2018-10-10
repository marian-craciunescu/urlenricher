// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/marian-craciunescu/urlenricher/metrics (interfaces: Endpoint)

// Package rest is a generated GoMock package.
package rest

import (
	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo"
	reflect "reflect"
)

// MockMetricsEndpoint is a mock of Endpoint interface
type MockMetricsEndpoint struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsEndpointMockRecorder
}

// MockMetricsEndpointMockRecorder is the mock recorder for MockMetricsEndpoint
type MockMetricsEndpointMockRecorder struct {
	mock *MockMetricsEndpoint
}

// NewMockMetricsEndpoint creates a new mock instance
func NewMockMetricsEndpoint(ctrl *gomock.Controller) *MockMetricsEndpoint {
	mock := &MockMetricsEndpoint{ctrl: ctrl}
	mock.recorder = &MockMetricsEndpointMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetricsEndpoint) EXPECT() *MockMetricsEndpointMockRecorder {
	return m.recorder
}

// Metrics mocks base method
func (m *MockMetricsEndpoint) Metrics(arg0 echo.Context) error {
	ret := m.ctrl.Call(m, "Metrics", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Metrics indicates an expected call of Metrics
func (mr *MockMetricsEndpointMockRecorder) Metrics(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metrics", reflect.TypeOf((*MockMetricsEndpoint)(nil).Metrics), arg0)
}