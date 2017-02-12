package strings

import (
	"github.com/geminikim/minimem/constant"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_SetByStringStoreManager(t *testing.T) {
	manager := NewStringStoreManager();

	value := make(map[string]string)
	value[constant.KEY] = "hello"
	value[constant.VALUE] = "HelloWorld"
	manager.Process(store.Message{constant.GET, value})

	response := manager.Process(store.Message{constant.SET, value})
	assert.Equal(t, constant.OK, response)
}

func Test_GetByStringStoreManager(t *testing.T) {
	manager := NewStringStoreManager();

	value := make(map[string]string)
	value[constant.KEY] = "hello"
	value[constant.VALUE] = "HelloWorld"
	manager.Process(store.Message{constant.SET, value})

	response := manager.Process(store.Message{constant.GET, value})
	assert.Equal(t, "HelloWorld", response)
}