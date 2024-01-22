package model

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewPattern(t *testing.T) {
	p := NewPattern()

	assert.Equal(t, "Rule", p.Name, "Expected Name to be 'Rule'")
	assert.NotEqual(t, uuid.Nil, p.Id, "Expected Id to be a non-zero UUID")
	assert.Equal(t, "", p.Before, "Expected Before to be an empty string")
	assert.Equal(t, "", p.After, "Expected After to be an empty string")
	assert.False(t, p.Replace, "Expected Replace to be false")
}
