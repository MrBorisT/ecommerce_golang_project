package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	RequestsCounter = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "checkout",
		Subsystem: "http",
		Name:      "requests_total",
	})
	ResponseCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "checkout",
		Subsystem: "http",
		Name:      "responses_total",
	},
		[]string{"status"},
	)
	HistogramResponseTime = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "checkout",
		Subsystem: "http",
		Name:      "histogram_response_time_seconds",
		Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 16),
	},
		[]string{"status"},
	)
	ProductServiceHistogramResponseTime = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "checkout",
		Subsystem: "product_service_grpc",
		Name:      "histogram_response_time_seconds",
		Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 16),
	},
		[]string{"status"},
	)
	LOMSHistogramResponseTime = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "checkout",
		Subsystem: "loms_grpc",
		Name:      "histogram_response_time_seconds",
		Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 16),
	},
		[]string{"status"},
	)
)

func New() http.Handler {
	return promhttp.Handler()
}
