package four

import (
	"time"

	"github.com/go-kit/kit/log"
	. "github.com/myesui/uuid"
	"github.com/myesui/uuid/kit"
)

var _ Service = &loggingMiddleware{}

type loggingMiddleware struct {
	kit.LoggingMiddleware
}

func (o loggingMiddleware) UUID() (id UUID) {
	defer func(begin time.Time) {
		o.Log(begin, "method", "uuid", "uuid", id)
	}(time.Now())

	id = o.Next().(Service).UUID()
	return
}

func (o loggingMiddleware) Bulk(amount int) (ids []UUID) {
	defer func(begin time.Time) {
		o.Log(begin, "method", "bulk", "amount", amount, "uuids", ids)
	}(time.Now())

	ids = o.Next().(Service).Bulk(amount)
	return
}

func (o *loggingMiddleware) Add(service kit.UUIDKit) Service {
	return kit.AddMiddleware(service, o).(Service)
}

func (loggingMiddleware) String() string {
	return "uuid-four-logging"
}

func NewLoggingMiddleware(logger log.Logger) Service {
	return &loggingMiddleware{
		kit.LoggingMiddleware{
			Logger: logger,
		},
	}
}
