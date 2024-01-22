package model

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewSequence(t *testing.T) {
	name := "TestSequence"
	content := "TestContent"

	sequence := NewSequence(name, content)

	assert.Equal(t, name, sequence.Name, "Expected Name to match input")
	assert.NotEqual(t, uuid.Nil, sequence.Id, "Expected Id to be a non-zero UUID")
	assert.Equal(t, content, sequence.Content, "Expected Content to match input")
	assert.NotNil(t, sequence.Precedents, "Expected Precedents to be a non-nil slice")
	assert.Empty(t, sequence.Precedents, "Expected Precedents to be an empty slice")
}
