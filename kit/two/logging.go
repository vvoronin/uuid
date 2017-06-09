package two

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

func (o loggingMiddleware) UUID(idType SystemId) (id UUID) {
	defer func(begin time.Time) {
		o.Log(begin, "method", "uuid", "type", idType.String(), "uuid", id, "error", "false")
	}(time.Now())

	id = o.Next().(Service).UUID(idType)
	return
}

func (o *loggingMiddleware) Add(service kit.UUIDKit) Service {
	return kit.AddMiddleware(service, o).(Service)
}

func (loggingMiddleware) String() string {
	return "uuid-two-logging"
}

func NewLoggingMiddleware(logger log.Logger) Service {
	return &loggingMiddleware{
		kit.LoggingMiddleware{
			Logger: logger,
		},
	}
}
