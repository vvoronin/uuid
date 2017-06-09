package five

import (
	. "github.com/myesui/uuid"
	"github.com/myesui/uuid/kit"
)

var _ Service = &service{}

type Service interface {
	kit.UUIDKit
	Add(kit.UUIDKit) Service

	UUID(Implementation, ...interface{}) UUID
}

type service struct {
	generator *Generator
	kit.Middleware
}

func (o service) UUID(namespace Implementation, names ...interface{}) UUID {
	return o.generator.NewV5(namespace, names...)
}

func (o *service) Add(service kit.UUIDKit) Service {
	return kit.AddMiddleware(service, o).(Service)
}

func (service) String() string {
	return "uuid-five"
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
