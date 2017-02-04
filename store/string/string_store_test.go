package strings

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_shouldBeSuccessSetAndGet(t *testing.T) {
	store := newStringStore()
	store.set("key", "value")
	value := store.get("key")

	assert.Equal(t, "value", value)
}