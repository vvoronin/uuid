package kit

import (
	"time"

	"github.com/go-kit/kit/log"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Middleware
}

func (o LoggingMiddleware) Log(begin time.Time, values ...interface{}) {
	o.Logger.Log(append(values, "took", time.Since(begin))...)
}
