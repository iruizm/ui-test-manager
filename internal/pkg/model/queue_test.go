package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	assert.NotNil(t, q)
	assert.Equal(t, 0, q.Size())
	assert.True(t, q.IsEmpty())
}

func TestEnqueueDequeue(t *testing.T) {
	q := NewQueue()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	item, ok := q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 1, item)
	assert.Equal(t, 2, q.Size())
	assert.False(t, q.IsEmpty())

	item, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 2, item)
	assert.Equal(t, 1, q.Size())
	assert.False(t, q.IsEmpty())

	item, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 3, item)
	assert.Equal(t, 0, q.Size())
	assert.True(t, q.IsEmpty())

	item, ok = q.Dequeue()
	assert.False(t, ok)
	assert.Nil(t, item)
}

func TestIsEmpty(t *testing.T) {
	q := NewQueue()
	assert.True(t, q.IsEmpty())

	q.Enqueue(1)
	assert.False(t, q.IsEmpty())

	q.Dequeue()
	assert.True(t, q.IsEmpty())
}

func TestSize(t *testing.T) {
	q := NewQueue()
	assert.Equal(t, 0, q.Size())

	q.Enqueue(1)
	assert.Equal(t, 1, q.Size())

	q.Enqueue(2)
	assert.Equal(t, 2, q.Size())

	q.Dequeue()
	assert.Equal(t, 1, q.Size())

	q.Dequeue()
	assert.Equal(t, 0, q.Size())
}
