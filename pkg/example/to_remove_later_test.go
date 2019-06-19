package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExampleOnSimpleConditions(t *testing.T) {
	assert.True(t, example(), "Example function returned true")
	assert.NotNil(t, example(), "Example not returned nil")
}
