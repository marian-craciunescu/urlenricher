// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/marian-craciunescu/urlenricher/connector (interfaces: Connector)

// Package cachestore is a generated GoMock package.
package cachestore

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/marian-craciunescu/urlenricher/models"
	reflect "reflect"
)

// MockConnector is a mock of Connector interface
type MockConnector struct {
	ctrl     *gomock.Controller
	recorder *MockConnectorMockRecorder
}

// MockConnectorMockRecorder is the mock recorder for MockConnector
type MockConnectorMockRecorder struct {
	mock *MockConnector
}

// NewMockConnector creates a new mock instance
func NewMockConnector(ctrl *gomock.Controller) *MockConnector {
	mock := &MockConnector{ctrl: ctrl}
	mock.recorder = &MockConnectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConnector) EXPECT() *MockConnectorMockRecorder {
	return m.recorder
}

// Resolve mocks base method
func (m *MockConnector) Resolve(arg0 string) (*models.URL, error) {
	ret := m.ctrl.Call(m, "Resolve", arg0)
	ret0, _ := ret[0].(*models.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resolve indicates an expected call of Resolve
func (mr *MockConnectorMockRecorder) Resolve(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockConnector)(nil).Resolve), arg0)
}

// Start mocks base method
func (m *MockConnector) Start() error {
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockConnectorMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockConnector)(nil).Start))
}

// Stop mocks base method
func (m *MockConnector) Stop() error {
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockConnectorMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockConnector)(nil).Stop))
}
