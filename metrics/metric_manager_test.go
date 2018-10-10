package metrics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestManagerSanity(t *testing.T) {
	a := assert.New(t)

	mgr := NewMetricManager()

	metric := mgr.RegisterMetric("test")
	intMetric := metric.AddInt("metric_name")
	intMetric.Add(1)

	m := mgr.FormatMetrics()
	a.Equal("1", m["test.metric_name"])

}
