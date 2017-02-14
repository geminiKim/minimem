package store


import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_shouldBeSuccessSetAndGetByHashStore(t *testing.T) {
	store := NewHashStore()
	store.Set("key", "field", "value")
	value := store.Get("key","field")

	assert.Equal(t, "value", value)
}