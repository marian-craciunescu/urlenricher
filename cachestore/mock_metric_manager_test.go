// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/marian-craciunescu/urlenricher/metrics (interfaces: MetricManager)

// Package cachestore is a generated GoMock package.
package cachestore

import (
	gomock "github.com/golang/mock/gomock"
	metrics "github.com/marian-craciunescu/urlenricher/metrics"
	reflect "reflect"
)

// MockMetricManager is a mock of MetricManager interface
type MockMetricManager struct {
	ctrl     *gomock.Controller
	recorder *MockMetricManagerMockRecorder
}

// MockMetricManagerMockRecorder is the mock recorder for MockMetricManager
type MockMetricManagerMockRecorder struct {
	mock *MockMetricManager
}

// NewMockMetricManager creates a new mock instance
func NewMockMetricManager(ctrl *gomock.Controller) *MockMetricManager {
	mock := &MockMetricManager{ctrl: ctrl}
	mock.recorder = &MockMetricManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetricManager) EXPECT() *MockMetricManagerMockRecorder {
	return m.recorder
}

// FormatMetrics mocks base method
func (m *MockMetricManager) FormatMetrics() map[string]interface{} {
	ret := m.ctrl.Call(m, "FormatMetrics")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// FormatMetrics indicates an expected call of FormatMetrics
func (mr *MockMetricManagerMockRecorder) FormatMetrics() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormatMetrics", reflect.TypeOf((*MockMetricManager)(nil).FormatMetrics))
}

// RegisterMetric mocks base method
func (m *MockMetricManager) RegisterMetric(arg0 string) *metrics.Metric {
	ret := m.ctrl.Call(m, "RegisterMetric", arg0)
	ret0, _ := ret[0].(*metrics.Metric)
	return ret0
}

// RegisterMetric indicates an expected call of RegisterMetric
func (mr *MockMetricManagerMockRecorder) RegisterMetric(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterMetric", reflect.TypeOf((*MockMetricManager)(nil).RegisterMetric), arg0)
}
