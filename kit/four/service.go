package four

import (
	. "github.com/myesui/uuid"
	"github.com/myesui/uuid/kit"
)

var _ Service = &service{}

type Service interface {
	kit.UUIDKit
	Add(kit.UUIDKit) Service

	UUID() UUID
	Bulk(int) []UUID
}

type service struct {
	generator *Generator
	kit.Middleware
}

func (o service) UUID() UUID {
	return o.generator.NewV4()
}

func (o service) Bulk(amount int) []UUID {
	return o.generator.BulkV4(amount)
}

func (o *service) Add(service kit.UUIDKit) Service {
	return kit.AddMiddleware(service, o).(Service)
}

func (service) String() string {
	return "uuid-one"
}

func NewService(config *GeneratorConfig) Service {
	generator, err := NewGenerator(config)
	if err != nil {
		panic(err)
	}
	return kit.Make(&service{
		generator: generator,
	}).(Service)
}
