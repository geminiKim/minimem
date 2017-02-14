package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/geminikim/minimem/util"
)

func Test_listStoreFunc(t *testing.T) {
	store := NewListStore()
	store.RightPush("key", "value3")
	store.RightPush("key", "value4")
	store.RightPush("key", "value5")
	store.LeftPush("key", "value2")
	store.LeftPush("key", "value1")
	store.LeftPush("key", "value0")


	assert.Equal(t, "value0", store.LeftPeek("key"))
	assert.Equal(t, "value5", store.RightPeek("key"))
	assert.Equal(t, "value0", store.LeftPop("key"))
	assert.Equal(t, "value5", store.RightPop("key"))

	list := store.ByRange("key", 0, 2)

	assert.Equal(t, "value1", list[0])
	assert.Equal(t, "value2", list[1])
}

func Test_listRangeGetToString(t *testing.T) {
	store := NewListStore()
	store.RightPush("key", "value3")
	store.RightPush("key", "value4")
	store.RightPush("key", "value5")
	store.LeftPush("key", "value2")
	store.LeftPush("key", "value1")
	store.LeftPush("key", "value0")
	assert.Equal(t, "[\"value0\",\"value1\"]", util.ToJsonString(store.ByRange("key", 0, 2)))

}