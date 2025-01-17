// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	metrics "github.com/rcrowley/go-metrics"

	mock "github.com/stretchr/testify/mock"

	sync "sync"

	time "time"
)

// MetricsManager is an autogenerated mock type for the MetricsManager type
type MetricsManager struct {
	mock.Mock
}

// GetCounter provides a mock function with given fields: name
func (_m *MetricsManager) GetCounter(name string) metrics.Counter {
	ret := _m.Called(name)

	var r0 metrics.Counter
	if rf, ok := ret.Get(0).(func(string) metrics.Counter); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metrics.Counter)
		}
	}

	return r0
}

// GetGauge provides a mock function with given fields: name
func (_m *MetricsManager) GetGauge(name string) metrics.Gauge {
	ret := _m.Called(name)

	var r0 metrics.Gauge
	if rf, ok := ret.Get(0).(func(string) metrics.Gauge); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metrics.Gauge)
		}
	}

	return r0
}

// GetGaugeFloat64 provides a mock function with given fields: name
func (_m *MetricsManager) GetGaugeFloat64(name string) metrics.GaugeFloat64 {
	ret := _m.Called(name)

	var r0 metrics.GaugeFloat64
	if rf, ok := ret.Get(0).(func(string) metrics.GaugeFloat64); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metrics.GaugeFloat64)
		}
	}

	return r0
}

// GetTimer provides a mock function with given fields: name
func (_m *MetricsManager) GetTimer(name string) metrics.Timer {
	ret := _m.Called(name)

	var r0 metrics.Timer
	if rf, ok := ret.Get(0).(func(string) metrics.Timer); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metrics.Timer)
		}
	}

	return r0
}

// Register provides a mock function with given fields: name, item, tags
func (_m *MetricsManager) Register(name string, item interface{}, tags map[string]string) error {
	ret := _m.Called(name, item, tags)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}, map[string]string) error); ok {
		r0 = rf(name, item, tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResetInterval provides a mock function with given fields: interval
func (_m *MetricsManager) ResetInterval(interval time.Duration) {
	_m.Called(interval)
}

// Run provides a mock function with given fields: ctx, wg
func (_m *MetricsManager) Run(ctx context.Context, wg *sync.WaitGroup) {
	_m.Called(ctx, wg)
}

// Unregister provides a mock function with given fields: name
func (_m *MetricsManager) Unregister(name string) {
	_m.Called(name)
}
