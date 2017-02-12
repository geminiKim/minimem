package strings

import (
	"github.com/geminikim/minimem/constant"
)

type StringStoreManager struct {
	store *stringStore
}

func NewStringStoreManager() store.Manager {
	manager := new(StringStoreManager)
	manager.store = newStringStore()
	return manager
}

func (manager StringStoreManager) Process(message store.Message) string {
	switch message.Command {
	case constant.GET: return manager.store.get(message.Value[constant.KEY])
	case constant.SET: return manager.store.set(message.Value[constant.KEY], message.Value[constant.VALUE])
	default: return constant.NOT_SUPPORTED_COMMAND
	}
}

func (manager StringStoreManager) GetType() string {
	return constant.STRING
}