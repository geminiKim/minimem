package hash


import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_shouldBeSuccessSetAndGetByHashStore(t *testing.T) {
	store := NewHashStore()
	store.set("key", "field", "value")
	value := store.get("key","field")

	assert.Equal(t, "value", value)
}