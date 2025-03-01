package k8soperators

// Go Prometheus metrics
import "github.com/prometheus/client_golang/prometheus"

var (
	requestsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total HTTP requests",
		},
		[]string{"path"},
	)
)
