package store

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_put(t *testing.T) {
	store := NewKeyValueStore()
	store.set("key", "value")
	value := store.get("key")

	assert.Equal(t, "value", value)
}