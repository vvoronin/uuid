package uuid_test

import (
	"gopkg.in/stretchr/testify.v1/assert"
	"io/ioutil"
	"log"
	"testing"

	. "github.com/myesui/uuid"
)

func TestInit(t *testing.T) {
	assert.Panics(t, didRegisterPanic, "Should panic")
}

func didRegisterPanic() {
	config := &GeneratorConfig{
		Logger: log.New(ioutil.Discard, "", 0),
	}
	RegisterGenerator(config)
	RegisterGenerator(config)
}
