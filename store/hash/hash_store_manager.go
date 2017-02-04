package hash

import (
	"github.com/geminikim/minimem/store"
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
	case "GET": return manager.store.get(message.Value["key"], message.Value["field"])
	case "SET": return manager.store.set(message.Value["key"], message.Value["field"], message.Value["value"])
	default: return "Not Supported Command"
	}
}

func (manager HashStoreManager) GetType() string {
	return "HASH"
}