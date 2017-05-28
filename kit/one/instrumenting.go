package one

import (
	. "github.com/myesui/uuid"
	"time"
	"github.com/myesui/uuid/kit"
)

var _ Service = &instrumentingMiddleware{}

type instrumentingMiddleware struct {
	kit.InstrumentingMiddleware
}

func (o instrumentingMiddleware) UUID() (id UUID) {
	defer func(begin time.Time) {
		o.Log(begin, "method", "uuid", "error", "false")
	}(time.Now())

	id = o.Next().(Service).UUID()
	return
}

func (o instrumentingMiddleware) Bulk(amount int) (ids []UUID) {
	defer func(begin time.Time) {
		o.Log(begin, "method", "bulk", "error", "false")
	}(time.Now())

	ids = o.Next().(Service).Bulk(amount)
	return
}

func (o *instrumentingMiddleware) Add(service kit.Kit) Service {
	return kit.AddMiddleware(service, o).(Service)
}

func (instrumentingMiddleware) String() string {
	return "uuid-one-instrumenting"
}

func NewInstrumentingMiddleware() kit.Kit {
	namespace := "uuid"
	subsystem := "one"

	return &instrumentingMiddleware{
		kit.NewInstrumentingMiddleware(namespace, subsystem),
	}
}
