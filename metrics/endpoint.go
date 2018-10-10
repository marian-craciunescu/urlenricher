package metrics

import (
	"github.com/labstack/echo"
	"net/http"
)

type Endpoint interface {
	Metrics(ctx echo.Context) error
}

type metricsEndpoint struct {
	m MetricManager
}

func NewMetricEnpoint(metric MetricManager) Endpoint {
	return &metricsEndpoint{metric}

}

func (e *metricsEndpoint) Metrics(ctx echo.Context) error {

	m := e.m.FormatMetrics()
	return ctx.JSON(http.StatusOK, m)
}
