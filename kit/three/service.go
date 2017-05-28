package three

import (
	. "github.com/myesui/uuid"
	"github.com/myesui/uuid/kit"
)

var  _ Service = &service{}

type Service interface {
	kit.Kit
	Add(kit.Kit) Service

	UUID(Implementation, ...interface{}) UUID
}

type service struct {
	generator *Generator
	kit.Middleware
}

func (o service) UUID(namespace Implementation, names ...interface{}) UUID {
	return o.generator.NewV3(namespace, names...)
}

func (o *service) Add(service kit.Kit) Service {
	return kit.AddMiddleware(service, o).(Service)
}

func (service) String() string {
	return "uuid-three"
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
