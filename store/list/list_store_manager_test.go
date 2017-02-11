package list

import (
	"github.com/geminikim/minimem/store"
	"github.com/geminikim/minimem/constant"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_LeftPushAndLeftPopByListStoreManager(t *testing.T) {
	manager := NewListStoreManager();

	value := make(map[string]string)
	value[constant.KEY] = "hello"
	value[constant.VALUE] = "HelloWorld"
	result := manager.Process(store.Message{constant.LEFT_PUSH, value})

	response := manager.Process(store.Message{constant.LEFT_POP, value})
	assert.Equal(t, constant.OK, result)
	assert.Equal(t, "HelloWorld", response)
}

func Test_LeftPushAndLeftPeekByListStoreManager(t *testing.T) {
	manager := NewListStoreManager();

	value := make(map[string]string)
	value[constant.KEY] = "hello"
	value[constant.VALUE] = "HelloWorld"
	result := manager.Process(store.Message{constant.LEFT_PUSH, value})

	response := manager.Process(store.Message{constant.LEFT_PEEK, value})
	assert.Equal(t, constant.OK, result)
	assert.Equal(t, "HelloWorld", response)
}

func Test_RightPushAndRightPopByListStoreManager(t *testing.T) {
	manager := NewListStoreManager();

	value := make(map[string]string)
	value[constant.KEY] = "hello"
	value[constant.VALUE] = "HelloWorld"
	result := manager.Process(store.Message{constant.RIGHT_PUSH, value})

	response := manager.Process(store.Message{constant.RIGHT_POP, value})
	assert.Equal(t, constant.OK, result)
	assert.Equal(t, "HelloWorld", response)
}

func Test_RightPushAndRightPeekByListStoreManager(t *testing.T) {
	manager := NewListStoreManager();

	value := make(map[string]string)
	value[constant.KEY] = "hello"
	value[constant.VALUE] = "HelloWorld"
	result := manager.Process(store.Message{constant.RIGHT_PUSH, value})

	response := manager.Process(store.Message{constant.RIGHT_PEEK, value})
	assert.Equal(t, constant.OK, result)
	assert.Equal(t, "HelloWorld", response)
}

func Test_ByRangeByListStoreManager(t *testing.T) {
	manager := NewListStoreManager();

	value := make(map[string]string)
	value[constant.KEY] = "hello"
	value[constant.VALUE] = "HelloWorld_0"
	manager.Process(store.Message{constant.RIGHT_PUSH, value})
	value[constant.VALUE] = "HelloWorld_1"
	manager.Process(store.Message{constant.RIGHT_PUSH, value})

	value[constant.KEY] = "hello"
	value[constant.INDEX] = "0"
	value[constant.COUNT] = "2"
	response := manager.Process(store.Message{constant.BY_RANGE, value})
	assert.Equal(t, "[\"HelloWorld_0\",\"HelloWorld_1\"]", response)
}