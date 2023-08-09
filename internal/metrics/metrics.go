package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	issueToken = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "issue_token_total",
			Help: "The total number of issue token",
		},
	)

	orderSuccess = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "order_success_total",
			Help: "The total number of order success",
		},
	)
)

func OrderSucceed() {
	orderSuccess.Inc()
}

func IssueToken() {
	issueToken.Inc()
}

type MetricStore struct{}

func (m MetricStore) ListenAndServe(addr string) error {
	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(addr, nil)
}

func RegisterMetrics() {
	prometheus.MustRegister(orderSuccess)
	prometheus.MustRegister(issueToken)
}
