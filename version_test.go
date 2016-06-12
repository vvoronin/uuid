package uuid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersion_String(t *testing.T) {
	for _, v := range []Version{
		One, Two, Three, Four, Five, Unknown,
	} {
		assert.NotEmpty(t, v.String(), "Expected a value")
	}
}
