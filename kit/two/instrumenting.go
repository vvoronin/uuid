package two

import (
	"time"

	. "github.com/myesui/uuid"
	"github.com/myesui/uuid/kit"
)

var _ Service = &instrumentingMiddleware{}

type instrumentingMiddleware struct {
	kit.InstrumentingMiddleware
}

func (o instrumentingMiddleware) UUID(idType SystemId) (id UUID) {
	defer func(begin time.Time) {
		o.Log(begin, "method", "uuid", "error", "false")
	}(time.Now())
	id = o.Next().(Service).UUID(idType)
	return
}

func (o *instrumentingMiddleware) Add(service kit.UUIDKit) Service {
	return kit.AddMiddleware(service, o).(Service)
}

func (instrumentingMiddleware) String() string {
	return "uuid-two-instrumenting"
}

func NewInstrumentingMiddleware() kit.UUIDKit {
	namespace := "uuid"
	subsystem := "two"

	return &instrumentingMiddleware{
		kit.NewInstrumentingMiddleware(namespace, subsystem),
	}
}
