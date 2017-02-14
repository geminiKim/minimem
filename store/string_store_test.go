package store

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_shouldBeSuccessSetAndGet(t *testing.T) {
	store := NewStringStore()
	store.Set("key", "value")
	value := store.Get("key")

	assert.Equal(t, "value", value)
}