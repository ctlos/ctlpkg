package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	c := BashCaller{}
	err := c.Execute("ls")
	assert.NoError(t, err)
}
