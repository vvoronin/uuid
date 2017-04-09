package uuid_test

import (
	. "github.com/twinj/uuid"
	"gopkg.in/stretchr/testify.v1/assert"
	"testing"
)

func TestInit(t *testing.T) {
	assert.Panics(t, didInitPanic, "Should panic")
}

func didInitPanic() {
	Init()
	Init()
}
