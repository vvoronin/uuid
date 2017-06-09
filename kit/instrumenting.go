package kit

import (
	"time"

	"github.com/go-kit/kit/metrics"
	kprometheus "github.com/go-kit/kit/metrics/prometheus"
	sprometheus "github.com/prometheus/client_golang/prometheus"
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
		requestCount: kprometheus.NewCounterFrom(sprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		requestLatency: kprometheus.NewSummaryFrom(sprometheus.SummaryOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
	}
}
