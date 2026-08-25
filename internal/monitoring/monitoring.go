package monitoring

import (
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var ready atomic.Bool

var (
	requestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "snooze",
			Subsystem: "http",
			Name:      "requests_total",
			Help:      "Total number of HTTP requests.",
		},
		[]string{"method", "route", "status"},
	)
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "snooze",
			Subsystem: "http",
			Name:      "request_duration_seconds",
			Help:      "HTTP request duration in seconds.",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"method", "route"},
	)
	inFlight = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "snooze",
			Subsystem: "http",
			Name:      "requests_in_flight",
			Help:      "Current number of HTTP requests in flight.",
		},
	)
)

func init() {
	prometheus.MustRegister(requestTotal, requestDuration, inFlight)
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		started := time.Now()
		inFlight.Inc()
		c.Next()
		inFlight.Dec()

		route := c.FullPath()
		if route == "" {
			route = "unmatched"
		}
		requestTotal.WithLabelValues(c.Request.Method, route, strconv.Itoa(c.Writer.Status())).Inc()
		requestDuration.WithLabelValues(c.Request.Method, route).Observe(time.Since(started).Seconds())
	}
}

func Routes(r *gin.Engine) {
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/healthz", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	r.GET("/readyz", func(c *gin.Context) {
		if !ready.Load() {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		c.Status(http.StatusOK)
	})
}

func SetReady(value bool) {
	ready.Store(value)
}
