package metrics

import (
	"expvar"
	"runtime"
	"time"
)

type MetricManager interface {
	FormatMetrics() map[string]interface{}
	RegisterMetric(name string) *Metric
}

type metricManager struct {
	registedMetric map[string]*Metric
}

func init() {
	start := time.Now()
	expvar.Publish("uptime", expvar.Func(func() interface{} {
		return time.Since(start).Seconds()
	}))
}

func NewMetricManager() MetricManager {
	m := make(map[string]*Metric, 0)
	return &metricManager{m}
}

func (mgr *metricManager) RegisterMetric(name string) *Metric {

	if val, ok := mgr.registedMetric[name]; ok {
		return val
	}
	n := newMetric(name)
	mgr.registedMetric[name] = n
	return n
}

func (mgr *metricManager) FormatMetrics() map[string]interface{} {
	memstatsFunc := expvar.Get("memstats").(expvar.Func)
	memstats := memstatsFunc().(runtime.MemStats)

	mm := make(map[string]interface{})
	mm["memstat"] = memstats
	mm["uptime"] = expvar.Get("uptime").String() + " s"
	for metricName := range mgr.registedMetric {

		metricMap := expvar.Get(metricName).(*expvar.Map)
		metricMap.Do(func(kv expvar.KeyValue) {
			mm[kv.Key] = kv.Value.String()
		})

	}

	return mm
}
