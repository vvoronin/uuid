package kit

import (
	"github.com/go-kit/kit/metrics"
	"time"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
)

type InstrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Middleware
}

func (o InstrumentingMiddleware) Log(begin time.Time, values ...string) {
	o.requestCount.With(values...).Add(1)
	o.requestLatency.With(values...).Observe(time.Since(begin).Seconds())
}

func NewInstrumentingMiddleware(namespace, subsystem string) InstrumentingMiddleware {
	fieldKeys := []string{"method", "error"}

	return InstrumentingMiddleware{
		requestCount: kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		requestLatency:kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
	}
}
