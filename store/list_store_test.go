package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_listStore(t *testing.T) {
	store := NewListStore()
	store.rightPush("key", "value3")
	store.rightPush("key", "value4")
	store.rightPush("key", "value5")
	store.leftPush("key", "value0")
	store.leftPush("key", "value1")
	store.leftPush("key", "value2")


	assert.Equal(t, "value0", store.leftPeek("key"))
	assert.Equal(t, "value5", store.rightPeek("key"))
	assert.Equal(t, "value0", store.leftPop("key"))
	assert.Equal(t, "value5", store.rightPop("key"))
	assert.Equal(t, "value2", store.indexPop("key", 1))
	assert.Equal(t, "value3", store.pop("key", "value3"))

	list := store.rangeGet("key", 0, 1)

	assert.Equal(t, "value1", list[0])
	assert.Equal(t, "value4", list[1])
}