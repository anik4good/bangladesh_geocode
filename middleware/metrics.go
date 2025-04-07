package middleware

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	}, []string{"method", "path", "status"})

	httpRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Duration of HTTP requests",
		Buckets: []float64{0.1, 0.3, 0.5, 0.7, 1, 3, 5, 7, 10},
	}, []string{"method", "path"})

	dbQueryDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "db_query_duration_seconds",
		Help:    "Duration of database queries",
		Buckets: []float64{0.01, 0.05, 0.1, 0.5, 1, 3, 5},
	}, []string{"query"})
)

func MetricsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		path := c.Path()
		method := c.Request().Method

		err := next(c)

		duration := time.Since(start).Seconds()
		status := c.Response().Status

		httpRequestsTotal.WithLabelValues(method, path, http.StatusText(status)).Inc()
		httpRequestDuration.WithLabelValues(method, path).Observe(duration)

		return err
	}
}

func RecordDBQueryDuration(query string, duration time.Duration) {
	dbQueryDuration.WithLabelValues(query).Observe(duration.Seconds())
}
