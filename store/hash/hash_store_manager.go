package hash

import (
	"github.com/geminikim/minimem/store"
	"github.com/geminikim/minimem/constant"
)

type HashStoreManager struct {
	store *hashStore
}

func NewHashStoreManager() store.Manager {
	manager := new(HashStoreManager)
	manager.store = NewHashStore()
	return manager
}

func (manager HashStoreManager) Process(message store.Message) string {
	switch message.Command {
	case constant.GET: return manager.store.get(message.Value[constant.KEY], message.Value[constant.FIELD])
	case constant.SET: return manager.store.set(message.Value[constant.KEY], message.Value[constant.FIELD], message.Value[constant.VALUE])
	default: return constant.NOT_SUPPORTED_COMMAND
	}
}

func (manager HashStoreManager) GetType() string {
	return constant.HASH
}