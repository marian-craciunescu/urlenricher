// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/marian-craciunescu/urlenricher/cachestore (interfaces: Endpoint)

// Package rest is a generated GoMock package.
package rest

import (
	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo"
	models "github.com/marian-craciunescu/urlenricher/models"
	reflect "reflect"
)

// MockCacheEndpoint is a mock of Endpoint interface
type MockCacheEndpoint struct {
	ctrl     *gomock.Controller
	recorder *MockCacheEndpointMockRecorder
}

// MockCacheEndpointMockRecorder is the mock recorder for MockCacheEndpoint
type MockCacheEndpointMockRecorder struct {
	mock *MockCacheEndpoint
}

// NewMockCacheEndpoint creates a new mock instance
func NewMockCacheEndpoint(ctrl *gomock.Controller) *MockCacheEndpoint {
	mock := &MockCacheEndpoint{ctrl: ctrl}
	mock.recorder = &MockCacheEndpointMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCacheEndpoint) EXPECT() *MockCacheEndpointMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockCacheEndpoint) Delete(arg0 string) error {
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockCacheEndpointMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCacheEndpoint)(nil).Delete), arg0)
}

// Get mocks base method
func (m *MockCacheEndpoint) Get(arg0 string) (*models.URL, error) {
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*models.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockCacheEndpointMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCacheEndpoint)(nil).Get), arg0)
}

// Resolve mocks base method
func (m *MockCacheEndpoint) Resolve(arg0 echo.Context) error {
	ret := m.ctrl.Call(m, "Resolve", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Resolve indicates an expected call of Resolve
func (mr *MockCacheEndpointMockRecorder) Resolve(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockCacheEndpoint)(nil).Resolve), arg0)
}

// Save mocks base method
func (m *MockCacheEndpoint) Save(arg0 *models.URL) error {
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockCacheEndpointMockRecorder) Save(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockCacheEndpoint)(nil).Save), arg0)
}