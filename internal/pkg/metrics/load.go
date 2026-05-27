package metrics

import "github.com/prometheus/client_golang/prometheus"

var HttpRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	},
	[]string{"path", "method", "status"},
)

var HttpRequestDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "Request duration",
	},
	[]string{"path", "method"},
)

func Init() {
	prometheus.MustRegister(HttpRequestDuration)
	prometheus.MustRegister(HttpRequestsTotal)
}
