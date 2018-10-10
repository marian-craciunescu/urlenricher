package metrics

import (
	"expvar"
)

var numGoroutines = expvar.NewInt("num_goroutines")

type Metric struct {
	Name string
	*expvar.Map
}

func newMetric(name string) *Metric {
	m := expvar.NewMap(name)
	return &Metric{name, m}
}

func (m *Metric) AddInt(name string) *expvar.Int {
	composedName := m.Name + "." + name
	n := expvar.NewInt(composedName)
	m.Set(composedName, n)
	return n
}

func (m *Metric) AddFloat(name string) *expvar.Float {
	n := expvar.NewFloat(name)
	m.Set(name, n)
	return n
}
