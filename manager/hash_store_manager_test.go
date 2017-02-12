package hash

import (
	"github.com/geminikim/minimem/constant"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_SetByHashStoreManager(t *testing.T) {
	manager := NewHashStoreManager();

	value := make(map[string]string)
	value[constant.KEY] = "hello"
	value[constant.FIELD] = "world"
	value[constant.VALUE] = "HelloWorld"
	manager.Process(store.Message{constant.GET, value})

	response := manager.Process(store.Message{constant.SET, value})
	assert.Equal(t, constant.OK, response)
}

func Test_GetByHashStoreManager(t *testing.T) {
	manager := NewHashStoreManager();

	value := make(map[string]string)
	value[constant.KEY] = "hello"
	value[constant.FIELD] = "world"
	value[constant.VALUE] = "HelloWorld"
	manager.Process(store.Message{constant.SET, value})

	response := manager.Process(store.Message{constant.GET, value})
	assert.Equal(t, "HelloWorld", response)
}